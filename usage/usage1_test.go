package usage

import (
	"testing"

	"github.com/cncsmonster/prioritydeque"
)

type Node struct {
	V        int
	InEdges  []*Edge
	OutEdges []*Edge
}

type Edge struct {
	V    int
	From *Node
	To   *Node
}

// 获取最小生成树
func findMinimunSpanningTree(edges []*Edge) []*Edge {
	out := []*Edge{}
	pdq := prioritydeque.New(func(x, y any) bool {
		return x.(*Edge).V < y.(*Edge).V
	})
	for _, edge := range edges {
		pdq.Push(edge)
	}
	passed := map[*Node]bool{}
	for pdq.Len() != 0 {
		min := pdq.PopMin().(*Edge)
		if passed[min.From] && passed[min.To] {
			continue
		}
		out = append(out, min)
		passed[min.From] = true
		passed[min.To] = true
	}
	return out
}

// 展示一种可能的用法

func TestMakeMinimunSpanningTree(t *testing.T) {
	// 构建一棵树,然后测试找到最小生成树
	nodes := make([]*Node, 5)
	for i := range nodes {
		nodes[i] = &Node{}
	}
	edges := make([]*Edge, 5)
	edges[0] = &Edge{1, nodes[0], nodes[1]}
	edges[1] = &Edge{2, nodes[0], nodes[2]}
	edges[2] = &Edge{3, nodes[1], nodes[3]}
	edges[3] = &Edge{4, nodes[3], nodes[4]}
	edges[4] = &Edge{5, nodes[2], nodes[3]}

	mmtree := findMinimunSpanningTree(edges)
	if len(mmtree) != 4 {
		t.Error(mmtree, len(mmtree))
	}
	has := map[*Edge]bool{
		edges[0]: true,
		edges[1]: true,
		edges[2]: true,
		edges[3]: true,
	}
	for _, e := range mmtree {
		if !has[e] {
			t.Errorf("%v should not contains %v", mmtree, e)
		}
	}
}
