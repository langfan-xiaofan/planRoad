package agent

import "C"
import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type Career struct {
	UserId         uint                   `json:"user_id"`
	Input          string                 `json:"input"`
	Messages       []*schema.Message      `json:"messages"`
	Persona        map[string]interface{} `json:"persona"`
	PersonaSummary map[string]interface{} `json:"personaSummary"`
	Intent         string                 `json:"intent"`
	Neo4jData      map[string]any         `json:"neo4jData"`
	IntentMessage  []*schema.Message      `json:"intentMessage"`
	Dao            dao.ConversationDao
}

func NewCareer(UserId uint, Input string, dao *dao.ConversationDao) *Career {
	return &Career{
		UserId:         UserId,
		Input:          Input,
		Messages:       make([]*schema.Message, 0),
		Persona:        make(map[string]interface{}),
		PersonaSummary: make(map[string]interface{}),
		Intent:         "",
		Neo4jData:      make(map[string]any),
		IntentMessage:  make([]*schema.Message, 0),
		Dao:            *dao,
	}
}

var (
	FileParser *openai.Client
	Agent      *react.Agent
	ChatModel  *eino.ChatModel
)

func Init() {
	FileParser = OpenAI()
	Agent = NewAgent()
	ChatModel = Eino()
}

func Eino() *eino.ChatModel {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	agentConfig := eino.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("Qwen_plusModel"),
		BaseURL: os.Getenv("BaseURL"),
	}
	client, err := eino.NewChatModel(context.Background(), &agentConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func CreateFileParseAgent() *eino.ChatModel {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	agentConfig := eino.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("Qwen_longModel"),
		BaseURL: os.Getenv("BaseURL"),
	}
	client, err := eino.NewChatModel(context.Background(), &agentConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func OpenAI() *openai.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config := openai.DefaultConfig(os.Getenv("API_KEY"))
	config.BaseURL = os.Getenv("BaseURL")
	client := openai.NewClientWithConfig(config)
	return client
}

func NewAgent() *react.Agent {
	client := Eino()
	agent, err := react.NewAgent(context.Background(), &react.AgentConfig{
		Model: client,
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return agent
}

func (c *Career) LoadContext(ctx context.Context, message []*schema.Message) ([]*schema.Message, error) {
	var err error
	var mess []*schema.Message
	mess, err = c.Dao.GetMessagesFromRedis(c.UserId)
	if err != nil {
		mess, err = c.Dao.GetMessagesFromMysql(c.UserId)
		if err != nil {
			return nil, err
		}
	}
	message = append(message, mess...)
	message = append(message, &schema.Message{Role: schema.User, Content: c.Input})
	fmt.Println(message)
	return message, nil
}

func IntentPrePare(ctx context.Context, in *Career) ([]*schema.Message, error) {
	return []*schema.Message{
		{
			Role:    schema.System,
			Content: IntentPrompt,
		},
		{
			Role:    schema.User,
			Content: in.Input,
		},
	}, nil
}

func (c *Career) SaveMessage(ctx context.Context, career *Career) (*Career, error) {
	var Message model.Message
	Message = model.Message{
		UserId:  career.UserId,
		Role:    career.Messages[len(career.Messages)-1].Role,
		Content: career.Messages[len(career.Messages)-1].Content,
	}
	err := career.Dao.AddMessageToRedis(career.UserId, &Message)
	if err != nil {
		return nil, err
	}
	err = career.Dao.AddMessageToMysql(&Message)
	return career, nil
}

func (c *Career) UpdatePersona(ctx context.Context, career *Career) (*Career, error) {
	picture, err := dao.GetUserPictureByUserId(c.UserId)
	if err != nil {
		return nil, err
	}
	message := make([]*schema.Message, 0)
	message = append(message, &schema.Message{
		Role:    schema.System,
		Content: ParseUserPicturePrompt,
	})
	message = append(message, c.Messages...)
	message = append(message, &schema.Message{
		Role:    schema.User,
		Content: "旧的用户画像如下：" + fmt.Sprint(picture) + "请你从上下文对话给出新的用户画像，如果没有需要更新的东西就不用更新，直接返回原本的用户画像",
	})
	res, err := ChatModel.Generate(ctx, message)
	if err != nil {
		return nil, err
	}
	var UserPicture dto.UserPictureReq
	err = json.Unmarshal([]byte(res.Content), &UserPicture)
	UserPicture.UserId = career.UserId
	err = dao.CreatUserPicture(UserPicture)
	if err != nil {
		return nil, err
	}
	return career, nil
}

func (c *Career) SummaryMessage(ctx context.Context, career *Career) (*Career, error) {
	career.Messages = append(career.Messages, &schema.Message{
		Role:    schema.User,
		Content: "总结一下对话",
	})
	generate, err := ChatModel.Generate(ctx, career.Messages)
	if err != nil {
		return nil, err
	}
	err = c.Dao.SummaryMessage(c.UserId, &[]model.Message{
		{
			Role:    schema.User,
			UserId:  career.UserId,
			Content: generate.Content,
		},
	})
	if err != nil {
		return career, err
	}
	return career, nil
}

func (c *Career) Graph(ctx context.Context, career *Career, client *eino.ChatModel, ch chan string) error {
	g := compose.NewGraph[*Career, *schema.Message](compose.WithGenLocalState(func(ctx context.Context) *Career {
		return career
	}))
	LoadHistoryMessageLambda := compose.InvokableLambda(c.LoadContext)
	IntentBranch := compose.NewGraphBranch(func(ctx context.Context, input *Career) (output string, err error) {
		if input.Intent == "career_plan" {
			fmt.Println("Intent is career_plan")
			return "career_plan", nil
		}
		if input.Intent == "career_switch" {
			fmt.Println("Intent is career_switch")
			return "career_switch", nil
		}
		if input.Intent == "promotion" {
			fmt.Println("Intent is promotion")
			return "promotion", nil
		}
		fmt.Println("Intent is general_chat")
		return "general_chat", nil
	}, map[string]bool{"career_plan": true, "career_switch": true, "promotion": true, "general_chat": true})
	PlanLambda := compose.InvokableLambda(func(ctx context.Context, input *Career) ([]*schema.Message, error) {
		return []*schema.Message{
			&schema.Message{
				Role:    schema.System,
				Content: PlanNodePrompt,
			},
		}, nil
	})
	SwitchLambda := compose.InvokableLambda(func(ctx context.Context, input *Career) ([]*schema.Message, error) {
		return []*schema.Message{
			&schema.Message{
				Role:    schema.System,
				Content: SwitchPrompt,
			},
		}, nil
	})
	PromotionLambda := compose.InvokableLambda(func(ctx context.Context, input *Career) ([]*schema.Message, error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: SwitchPrompt,
			},
		}, nil
	})
	GeneralLambda := compose.InvokableLambda(func(ctx context.Context, input *Career) ([]*schema.Message, error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: ChatPrompt,
			},
		}, nil
	})
	err := g.AddLambdaNode("intent_prepare", compose.InvokableLambda(IntentPrePare))
	if err != nil {
		return err
	}
	err = g.AddChatModelNode("intent_classifier", client, compose.WithStatePostHandler(func(ctx context.Context, out *schema.Message, state *Career) (*schema.Message, error) {
		state.Intent = out.Content
		return out, nil
	}))
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("intent_pos", compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output *Career, err error) {
		c.Intent = input.Content
		return c, nil
	}))
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("load_context", LoadHistoryMessageLambda)
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("career_plan", PlanLambda)
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("career_switch", SwitchLambda)
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("promotion", PromotionLambda)
	if err != nil {
		return err
	}
	err = g.AddLambdaNode("general_chat", GeneralLambda)
	if err != nil {
		return err
	}
	err = g.AddChatModelNode("chat", client)
	if err != nil {
		return err
	}
	err = g.AddBranch("intent_pos", IntentBranch)
	if err != nil {
		return err
	}
	err = g.AddEdge(compose.START, "intent_prepare")
	if err != nil {
		return err
	}
	err = g.AddEdge("intent_prepare", "intent_classifier")
	if err != nil {
		return err
	}
	err = g.AddEdge("intent_classifier", "intent_pos")
	if err != nil {
		return err
	}
	err = g.AddEdge("career_plan", "load_context")
	if err != nil {
		return err
	}
	err = g.AddEdge("promotion", "load_context")
	if err != nil {
		return err
	}
	err = g.AddEdge("general_chat", "load_context")
	if err != nil {
		return err
	}
	err = g.AddEdge("career_switch", "load_context")
	if err != nil {
		return err
	}
	err = g.AddEdge("load_context", "chat")
	if err != nil {
		return err
	}
	err = g.AddEdge("chat", compose.END)
	if err != nil {
		return err
	}
	r, err := g.Compile(ctx)
	if err != nil {
		return err
	}
	var content string
	result, err := r.Stream(ctx, career)
	if err != nil {
		return err
	}
	for {
		mess, err := result.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		}
		fmt.Print(mess.Content)
		ch <- mess.Content
		content += mess.Content
	}
	err = career.Dao.AddMessageToRedis(career.UserId, &model.Message{
		UserId:  career.UserId,
		Content: career.Input,
		Role:    schema.User,
	})
	err = career.Dao.AddMessageToRedis(career.UserId, &model.Message{
		UserId:  career.UserId,
		Role:    schema.Assistant,
		Content: content,
	})
	err = career.Dao.AddMessageToMysql(&model.Message{
		UserId:  career.UserId,
		Role:    schema.User,
		Content: career.Input,
	})
	err = career.Dao.AddMessageToMysql(&model.Message{
		UserId:  career.UserId,
		Role:    schema.Assistant,
		Content: content,
	})
	career.Messages = append(career.Messages, &schema.Message{
		Role:    schema.User,
		Content: career.Input,
	})
	career.Messages = append(career.Messages, &schema.Message{
		Role:    schema.Assistant,
		Content: content,
	})
	if len(career.Messages) > 20 {
		_, err := career.SummaryMessage(ctx, c)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func PurposeChain(input string, ch chan string, client *eino.ChatModel) (string, error) {
	ctx := context.Background()
	chain := compose.NewChain[string, *schema.Message]()
	chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input string) ([]*schema.Message, error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个根据用户的输入分析用户的意图的助手，你需要根据用户的输入，分析出用户的意图。仅仅返回用户的意图，只需要返回以下几个意图:学习规划,晋升路径,转型,其他",
			},
			{
				Role:    schema.User,
				Content: input,
			},
		}, nil
	})).AppendChatModel(client)
	c, err := chain.Compile(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	output, err := c.Stream(ctx, input)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var content string
	for {
		mess, err := output.Recv()
		if err != nil {
			if err == io.EOF {
				close(ch)
				break
			} else {
				fmt.Println(err)
			}
		}
		content += mess.Content
		ch <- mess.Content
	}
	return content, nil
}

func GenerateResumeInsight(data map[string]interface{}, ch chan string, client *eino.ChatModel) error {
	template := prompt.FromMessages(schema.GoTemplate,
		schema.SystemMessage(ParsePrompt))

	messages, err := template.Format(context.Background(), data)
	if err != nil {
		return err
	}
	stream, err := client.Stream(context.Background(), messages)
	if err != nil {
		return err
	}
	for {
		mess, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				close(ch)
				break
			}
			return err
		}
		ch <- mess.Content
	}
	return nil
}

func GenerateReportInsight(data map[string]interface{}, ch chan string, client *eino.ChatModel) error {
	return nil
}

func GenerateResumeRadar(data map[string]interface{}, client *eino.ChatModel) (dto.ResumeRadarRes, error) {
	var radar dto.ResumeRadarRes
	template := prompt.FromMessages(schema.GoTemplate, schema.SystemMessage(ResumeRadarPrompt))
	messages, err := template.Format(context.Background(), data)
	if err != nil {
		return dto.ResumeRadarRes{}, err
	}
	result, err := client.Generate(context.Background(), messages)
	if err != nil {
		return dto.ResumeRadarRes{}, err
	}
	err = json.Unmarshal([]byte(result.Content), &radar)
	if err != nil {
		return dto.ResumeRadarRes{}, err
	}
	return radar, nil
}
