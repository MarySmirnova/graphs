package graph

import "fmt"

//BFS реализует поиск вширину без учета весов
func (g *Graph) BFS(val string) (*Node, error) {
	var q []*Node
	el := g.Nodes[0]
	visit := make(map[*Node]bool)
	visit[el] = true

	for {
		if el.Value == val {
			return el, nil
		}
		q = g.queue(el, q, visit)
		if q == nil {
			return nil, fmt.Errorf("Элемент не найден")
		}
		el, q = dequeue(q)
	}
}

func dequeue(q []*Node) (*Node, []*Node) {
	if len(q) < 2 {
		return q[0], nil
	}
	return q[0], q[1:]
}

func (g Graph) queue(n *Node, q []*Node, visit map[*Node]bool) []*Node {
	for i := range n.Matrix {
		if n.Matrix[i] != 0 && !visit[g.Nodes[i]] {
			q = append(q, g.Nodes[i])
			visit[g.Nodes[i]] = true
		}
	}
	return q
}
