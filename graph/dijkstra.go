package graph

import (
	"fmt"
	"math"
)

//Dijkstra определяет кратчайший путь
//между любыми двумя узлами с учетом весов
func (g Graph) Dijkstra(start, end *Node) ([]*Node, error) {
	costs := make(map[*Node]float64)
	parents := make(map[*Node]*Node)
	visit := make(map[*Node]bool)
	begin := start
	var way []*Node

	//заполняем мапы начальными данными
	for i := range g.Nodes {
		costs[g.Nodes[i]] = math.Inf(1)
		parents[g.Nodes[i]] = nil
		visit[g.Nodes[i]] = false

		if g.Nodes[i] == begin {
			costs[begin] = 0
			parents[begin] = begin
		}
	}

	//ищем путь с наименьшей стоимостью до каждого узла
	for {
		for i := range begin.Matrix {
			if begin.Matrix[i] > 0 && begin.Matrix[i]+costs[begin] < costs[g.Nodes[i]] {
				costs[g.Nodes[i]] = begin.Matrix[i] + costs[begin]
				parents[g.Nodes[i]] = begin
			}
		}
		visit[begin] = true
		begin = findMinCost(costs, visit)
		if begin == nil {
			break
		}
	}

	//строим путь до заданного узла
	if !visit[end] {
		return nil, fmt.Errorf("Пути от узла %v до узла %v не существует", start.Value, end.Value)
	}
	for {
		way = append(way, end)
		if end == start {
			break
		}
		end = parents[end]
	}
	for i, j := 0, len(way)-1; i < j; i, j = i+1, j-1 {
		way[i], way[j] = way[j], way[i]
	}
	return way, nil
}

//indMinCost находит необработанный узел с наименьшей стоимостью
func findMinCost(costs map[*Node]float64, visit map[*Node]bool) *Node {
	minCost := math.Inf(1)
	var minNode *Node = nil
	for i := range costs {
		if costs[i] < minCost && !visit[i] {
			minCost = costs[i]
			minNode = i
		}
	}
	return minNode
}
