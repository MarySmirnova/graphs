package graph

//Структуры и конструктор взвешенного графа,
//Методы добавления узлов и ребер

type Node struct {
	Value  string
	ID     int
	Matrix []float64
}

type Graph struct {
	Nodes []*Node
	Count int
}

//NewGraph создает новый граф
func NewGraph() Graph {
	return Graph{nil, 0}
}

//NewNode создает новый узел, имеющий уникальный id
func (g *Graph) NewNode(val string) *Node {
	g.Count++
	n := &Node{val, g.Count, make([]float64, g.Count)}
	for i := range g.Nodes {
		g.Nodes[i].Matrix = append(g.Nodes[i].Matrix, 0)
	}
	g.Nodes = append(g.Nodes, n)
	return n
}

//NewEdge добавляет направленное взвешенное ребро
func (g *Graph) NewEdge(n1, n2 *Node, w float64) {
	n1.Matrix[n2.ID-1] = w
}

//AddEdgeUndir добавляет ненаправленное взвешеное ребро
func (g *Graph) NewEdgeUndir(n1, n2 *Node, w float64) {
	g.NewEdge(n1, n2, w)
	g.NewEdge(n2, n1, w)
}
