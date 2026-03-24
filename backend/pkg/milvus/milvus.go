package milvus

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

type MilvusConfig struct {
	Host string
	Port int
}

type Milvus struct {
	Client client.Client
	Config MilvusConfig
	Schema entity.Schema
}

func (m *Milvus) Init() error {
	ctx, err := context.WithTimeout(context.Background(), time.Second*5)
	if err != nil {
	}
	c, _ := client.NewGrpcClient(ctx, m.Config.Host+":"+strconv.Itoa(m.Config.Port))
	m.Client = c
	return nil
}

func (m *Milvus) Connect(collectionName string, schema entity.Schema) error {
	err := m.Client.LoadCollection(context.Background(), collectionName, false)
	if err != nil {
		return err
	}
	m.Schema = schema
	return nil
}

func (m *Milvus) Insert(collectionName []string, columns []entity.Column) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//for i := range columns {
	//	_, err := m.Client.Insert(ctx, collectionName[i], "", columns[i])
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//}
	columns = columns[1:]
	_, err = m.Client.Insert(ctx, collectionName[0], "", columns...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Milvus) MakeJobColumns(jobInfo interface{}) []entity.Column {
	var columns []entity.Column
	v := reflect.ValueOf(jobInfo)
	t := reflect.TypeOf(jobInfo)
	if t.Kind() != reflect.TypeOf(dao.JobInfo{}).Kind() {
		// fmt.Println("jobInfo is not dao.JobInfo type")
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		// fmt.Println(t.Field(i).Tag.Get("json"), v.Field(i))
		if v.Field(i).Kind() == reflect.Int64 {
			columns = append(columns, entity.NewColumnInt64(t.Field(i).Tag.Get("json"), []int64{v.Field(i).Int()}))
		} else if v.Field(i).Kind() == reflect.Array || v.Field(i).Kind() == reflect.Slice {
			if t.Field(i).Type.Elem().Kind() == reflect.Float32 {
				if v.Field(i).Len() == 0 {
					columns = append(columns, entity.NewColumnFloatVector(t.Field(i).Tag.Get("json"), 1024, make([][]float32, 1024)))
				} else {
					columns = append(columns, entity.NewColumnFloatVector(t.Field(i).Tag.Get("json"), 1024, [][]float32{v.Field(i).Interface().([]float32)}))
				}
			} else if t.Field(i).Type.Elem().Kind() == reflect.String {
				if v.Field(i).Len() == 0 {
					columns = append(columns, entity.NewColumnVarCharArray(t.Field(i).Tag.Get("json"), [][][]byte{{[]byte("")}}))
				} else {
					// 将[]string转换为[][]byte
					byteSlice := make([][]byte, len(v.Field(i).Interface().([]string)))
					for j, s := range v.Field(i).Interface().([]string) {
						byteSlice[j] = []byte(s)
					}
					columns = append(columns, entity.NewColumnVarCharArray(t.Field(i).Tag.Get("json"), [][][]byte{byteSlice}))
				}
				// // Handle string slice as Milvus Array type
				// var stringSlice []string
				// if v.Field(i).Len() > 0 {
				// 	stringSlice = v.Field(i).Interface().([]string)
				// } else {
				// 	stringSlice = []string{}
				// }
				// // Create array column for Milvus Array type
				// // 将[]string转换为[][]byte
				// byteSlice := make([][]byte, len(stringSlice))
				// for j, s := range stringSlice {
				// 	byteSlice[j] = []byte(s)
				// }
				// columns = append(columns, entity.NewColumnVarCharArray(t.Field(i).Tag.Get("json"), [][][]byte{byteSlice}))
			}
		} else if v.Field(i).Kind() == reflect.String {
			columns = append(columns, entity.NewColumnVarChar(t.Field(i).Tag.Get("json"), []string{v.Field(i).String()}))
		} else if v.Field(i).Kind() == reflect.Bool {
			columns = append(columns, entity.NewColumnBool(t.Field(i).Tag.Get("json"), []bool{v.Field(i).Bool()}))
		}
	}
	// fmt.Println("columns:", columns)
	return columns
}

func (m *Milvus) MakeOriginJob(jobinfo dto.JobInfo, vector []float32, id int64) dao.JobInfo {
	return dao.JobInfo{
		JobID:        id,
		WorkCity:     jobinfo.BasicInfo.WorkCity,
		PositionType: jobinfo.BasicInfo.PositionType,
		MinEducation: jobinfo.Requirement.Education.MinEducation,
		CoreSkills:   jobinfo.Requirement.SkillRequire.CoreSkills,
		Priority:     jobinfo.Priority.PreferSkills,
		SalaryRange:  jobinfo.WelfareInfo.SalaryRange,
		CompanyName:  jobinfo.BasicInfo.CompanyName,
		AcceptIntern: jobinfo.Requirement.Education.AcceptIntern == 1,
		NeedWorkExp:  jobinfo.Requirement.Experience.NeedWorkExp,
		Industry:     jobinfo.CompanyInfo.Industry,
		JobVector:    dao.JobVector{SkillVector: vector},
		PositionName: jobinfo.BasicInfo.PositionName,
	}
}

func (m *Milvus) Search(collectionName string, vector []float32, topK int64, sp entity.SearchParam) ([]client.SearchResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	results, err := m.Client.Search(ctx, collectionName, nil, "", []string{"position_name", "company_name", "work_city"}, []entity.Vector{entity.FloatVector(vector)}, "job_embedding", entity.COSINE, int(topK), sp)
	if err != nil {
		return nil, err
	}
	// for _, result := range results[0].Fields {
	// 	fmt.Println(result.FieldData().GetScalars().GetStringData().Data)
	// 	// for _, positionName := range result.FieldData().GetScalars().GetStringData().Data {
	// 	// 	fmt.Println(positionName)
	// 	// }
	// }
	return results, nil
}
