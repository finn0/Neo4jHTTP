package neo4j

import (
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
