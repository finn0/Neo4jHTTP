package neo4j

import (
	"fmt"
	"testing"
)

func TestMGraph(t *testing.T) {
	mGraph := &MGraph{}
	mGraph.CreateALGraph()

	for i, v := range mGraph.AdjList {
		fmt.Printf("index: %d, vertex: %s", i, v.Data)
		if v.FirstEdge != nil {
			print(v.FirstEdge)
		} else {
			fmt.Println()
		}
	}
}

func print(e *EdgeNode) {
	if e != nil {
		fmt.Printf(" -> %d:%v", e.Index, e.StockPercent)
	}
	if e.Next != nil {
		print(e.Next)
	} else {
		fmt.Printf("\n")
	}
}
