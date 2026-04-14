package dao

import (
	"backend/internal/db"
	"backend/internal/model"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func FindAllPaths(node string) (*model.SubgraphResponse, error) {
	ctx := context.Background()
	if db.Neo4jDb == nil {
		return nil, fmt.Errorf("neo4j database not initialized")
	}

	const cypher = `
       MATCH p1 = (p:position {name: $name})<-[:HAS_POSITION]-(j:job)-[:REQUIRES_BASE]->(base)
       MATCH p2 = (p)-[*]->(m)
       RETURN p1, p2
    `

	session := db.Neo4jDb.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, cypher, map[string]any{"name": node})
		if err != nil {
			return nil, err
		}

		nodeMap := make(map[string]model.Node)        // Key 改为 string 类型的 ElementId
		relMap := make(map[string]model.Relationship) // Key 改为 string 类型的 ElementId

		for res.Next(ctx) {
			for _, val := range res.Record().Values {
				if path, ok := val.(neo4j.Path); ok {
					// 1. 处理点
					for _, n := range path.Nodes {
						props := n.GetProperties()
						name, _ := props["name"].(string) // 尝试提取 name 属性

						nodeMap[n.ElementId] = model.Node{
							ID:    n.ElementId,
							Label: n.Labels[0], // 取第一个 Label 作为分类
							Name:  name,        // 前端直接用来显示在圆圈上
							//							Properties: props,
						}
					}
					// 2. 处理边
					for _, rel := range path.Relationships {
						relMap[rel.ElementId] = model.Relationship{
							ID:     rel.ElementId,
							Source: rel.StartElementId, // 对应前端的 source
							Target: rel.EndElementId,   // 对应前端的 target
							Type:   rel.Type,
						}
					}
				}
			}
		}

		// 3. 转为切片格式返回
		resp := &model.SubgraphResponse{
			Nodes:         make([]model.Node, 0, len(nodeMap)),
			Relationships: make([]model.Relationship, 0, len(relMap)),
		}
		for _, n := range nodeMap {
			resp.Nodes = append(resp.Nodes, n)
		}
		for _, r := range relMap {
			resp.Relationships = append(resp.Relationships, r)
		}
		return resp, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*model.SubgraphResponse), nil
}

//MATCH (p:position {name: "Go 微服务与云原生方向"})<-[:HAS_POSITION]-(j:job)-[:REQUIRES_BASE]->(base)
//RETURN base
//MATCH p = (start:position {name: "Go 微服务与云原生方向"})-[*]->(m)
//RETURN p
