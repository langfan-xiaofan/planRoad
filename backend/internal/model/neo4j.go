package model

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"` // 取第一个标签即可
	Name  string `json:"name"`  // 提取核心显示文字
	//Properties map[string]any `json:"properties"` // 备用属性
}

type Relationship struct {
	ID     string `json:"id"`
	Source string `json:"source"` // 前端库通用字段名
	Target string `json:"target"` // 前端库通用字段名
	Type   string `json:"type"`
}

type SubgraphResponse struct {
	Nodes         []Node         `json:"nodes"`
	Relationships []Relationship `json:"relationships"`
}
