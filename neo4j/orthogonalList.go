package neo4j

// VertexNode -
// data firstin firstout
type VertexNode struct {
	index    int
	FirstIn  *EdgeNode
	FirstOut *EdgeNode
}

// EdgeNode -
// tailvex headvex headlink taillink
type EdgeNode struct {
	TailVex  int
	HeadVex  int
	HeadLink *EdgeNode
	TailLink *EdgeNode
}

// OrthogonalList -
type OrthogonalList struct {
	VexList  []*VertexNode
	numNodes int
	numRels  int
}

// net
// 0->3
// 1->0, 1->2
// 2<->1
// 3

// matrix
// var adjMatrix = [][]int{
// 	{0, 0, 0, 1},
// 	{1, 0, 1, 0},
// 	{1, 1, 0, 0},
// 	{0, 0, 0, 0},
// }

var nodes = []int{0, 1, 2, 3}
var outRels = [][]int{
	{3},    // 0
	{0, 2}, // 1
	{1, 0}, // 2
	{},     // 3
}
var inRels = [][]int{
	{1, 2}, // 0
	{2},    // 1
	{1},    // 2
	{0},    // 3
}

// CreateGraph -
func CreateGraph() *OrthogonalList {
	var l OrthogonalList
	l.numNodes = len(nodes)
	for _, s := range outRels {
		l.numRels += len(s)
	}

	// 邻接表
	for i, es := range outRels {
		// 顶点表
		v := &VertexNode{
			index: i,
		}

		// 查找行表尾节点, 并设置新的尾节点
		for _, j := range es {
			if v.FirstOut == nil {
				v.FirstOut = &EdgeNode{
					TailVex: i,
					HeadVex: j,
				}
			} else {
				v.outTail().TailLink = &EdgeNode{
					TailVex: i,
					HeadVex: j,
				}
			}
		}
		l.VexList = append(l.VexList, v)
	}

	// 逆邻接表
	for i, es := range inRels {
		// 顶点
		v := l.VexList[i]

		for _, j := range es {
			e := l.VexList[j].findEdge(j, i)
			if v.FirstIn == nil {
				v.FirstIn = e
			} else {
				v.inTail().HeadLink = e
			}
		}
	}

	return &l
}

// 获取出度尾节点
func (v *VertexNode) outTail() *EdgeNode {
	if v.FirstOut == nil {
		return nil
	}
	t := v.FirstOut
	for t.TailLink != nil {
		t = t.TailLink
	}
	return t
}

// 获取入度尾节点
func (v *VertexNode) inTail() *EdgeNode {
	if v.FirstIn == nil {
		return nil
	}
	t := v.FirstIn
	for t.HeadLink != nil {
		t = t.HeadLink
	}
	return t
}

func (v *VertexNode) findEdge(tailVex, headVex int) *EdgeNode {
	var want *EdgeNode
	t := v.FirstOut
	for t != nil {
		if t.TailVex == tailVex && t.HeadVex == headVex {
			want = t
			break
		}
		t = t.TailLink
	}
	return want
}

// DFS -
func (l *OrthogonalList) DFS(i int, visited *map[int]bool) {
	if !(*visited)[i] {
		(*visited)[i] = true
	}

	e := l.VexList[i].FirstOut
	for e != nil {
		if !(*visited)[e.HeadVex] {
			l.DFS(e.HeadVex, visited)
		}
		e = e.TailLink
	}
}

// DFSTraverse -
func (l *OrthogonalList) DFSTraverse(i int, visited *map[int]bool) {
	if !(*visited)[i] {
		(*visited)[i] = true
	}

	e := l.VexList[i].FirstIn
	for e != nil {
		if !(*visited)[e.TailVex] {
			l.DFSTraverse(e.TailVex, visited)
		}
		e = e.HeadLink
	}
}
