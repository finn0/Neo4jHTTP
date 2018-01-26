package neo4j

import (
	"strconv"
)

// 十字链表，有向图
// 出度 + 入度

type Vertex struct {
	Data    string
	OutLink *Edge // 出度边表头指针
	InLink  *Edge // 入度边表头指针
}

type Edge struct {
	SourceIndex int    // 起点下标
	TargetIndex int    // 终点下标
	Data        string // 边属性
	InNext      *Edge  // 下一个逆邻接点
	OutNext     *Edge  // 下一个邻接点
}

type ArcGraph struct {
	AdjList   []Vertex
	NumVertex int
	NumEdge   int
}

func (g *ArcGraph) CreateGraph() *ArcGraph {
	g.NumVertex = 4
	g.NumEdge = 5

	// gen vertex node list
	vertexList := make([]Vertex, g.NumVertex)
	for i := 0; i < g.NumVertex; i++ {
		vertexList[i].Data = "V" + strconv.Itoa(i)
	}
	g.AdjList = vertexList

	// gen adjList
	// 0
	g.AdjList[0].OutLink = &Edge{0, 3, "0->3", nil, nil}
	g.AdjList[0].InLink = &Edge{1, 0, "1->0", &Edge{2, 0, "2->0", nil, nil}, nil}
	// 1
	v1 := g.AdjList[1]
	e := findEdgeByInLink(&v1, 1, 0)
	if e != nil {
		v1.OutLink = e
	} else {
		v1.OutLink = &Edge{1, 0, "1->0", nil, nil}
	}
	e = findEdgeByInLink(&v1, 1, 2)
	if e != nil {
		v1.OutLink.OutNext = e
	} else {
		v1.OutLink.OutNext = &Edge{1, 2, "1->2", nil, nil}
	}
	// 2
	// to be continued

	return g
}

func findEdgeByInLink(v *Vertex, source, target int) *Edge {
	if v.InLink == nil {
		return nil
	}
	if v.InLink.SourceIndex == source && v.InLink.TargetIndex == target {
		return v.InLink
	}
	return nil
}
