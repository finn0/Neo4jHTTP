package neo4j

import (
	"fmt"
	"testing"
)

func TestOrthogonalList(t *testing.T) {
	l := CreateGraph()

	if l.VexList[0].FirstOut != l.VexList[3].FirstIn ||
		l.VexList[1].FirstOut != l.VexList[0].FirstIn ||
		l.VexList[1].FirstOut.TailLink != l.VexList[2].FirstIn ||
		l.VexList[0].FirstIn.HeadLink != l.VexList[2].FirstOut.TailLink {
		t.Error("Wrong structure")
	}
}

func TestOrthogonalListDFS(t *testing.T) {
	visited := make(map[int]bool)
	l := CreateGraph()

	for i := 0; i < l.numNodes; i++ {
		l.DFS(i, &visited)
	}

	fmt.Printf("%v\n", visited)
}

func TestOrthogonalListDFSTraverse(t *testing.T) {
	visited := make(map[int]bool)
	l := CreateGraph()

	for i := 0; i < l.numNodes; i++ {
		l.DFSTraverse(i, &visited)
	}
	fmt.Printf("%v\n", visited)
}
