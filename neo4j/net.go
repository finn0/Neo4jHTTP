package neo4j

import (
	"fmt"
	"strconv"
)

// 有向网图
// 邻接表表示法，出度友好的, 出度 = adjList.size - 1

// === RUN   TestMGraph
// index: 0, vertex: V0 -> 4:6
// index: 1, vertex: V1 -> 0:9 -> 2:3
// index: 2, vertex: V2 -> 3:5 -> 0:2
// index: 3, vertex: V3 -> 4:1
// index: 4, vertex: V4
// --- PASS: TestMGraph (0.00s)

// VertexType - 顶点
type VertexType string

// EdgeType - 边
type EdgeType int

// VertexNode - 顶表节点
type VertexNode struct {
	Data      VertexType // 顶点域，存储顶点信息
	FirstEdge *EdgeNode  // 边表头指针
}

// EdgeNode - 边表节点
type EdgeNode struct {
	Index        int       // 邻接点域，存储该节点的下标
	StockPercent EdgeType  // 出度关系的数据 rel.properties
	Next         *EdgeNode // 邻域，指向下一个邻接点
}

type MGraph struct {
	AdjList     []VertexNode // 节点slice
	NumVertexes int          // 网图中节点数
	NumEdges    int          // 网图中边数
}

func (g *MGraph) CreateALGraph() *MGraph {
	g.NumVertexes = 5
	g.NumEdges = 6

	// gen vertex node list
	adjList := make([]VertexNode, g.NumVertexes)
	for i := 0; i < g.NumVertexes; i++ {
		adjList[i] = VertexNode{
			Data: VertexType("V" + strconv.Itoa(i)),
		}
	}
	g.AdjList = adjList

	// gen adjacency list
	// 0->4
	g.AdjList[0].FirstEdge = &EdgeNode{4, 6, nil}
	g.AdjList[1].FirstEdge = &EdgeNode{0, 9, &EdgeNode{2, 3, nil}}
	g.AdjList[2].FirstEdge = &EdgeNode{3, 5, &EdgeNode{0, 2, nil}}
	g.AdjList[3].FirstEdge = &EdgeNode{4, 1, nil}
	g.AdjList[4].FirstEdge = nil

	return g
}

func scanf(str string) int64 {
	fmt.Print(str)
	var s string
	fmt.Scanf("%s", &s)
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// VertexNode ------ EdgeNode
// [V0, ] ---------- [V4, Data]-[V5, Data]-[...]
// [V1, ] ---------- []
// [V2, ] ---------- []
// ...
