package main

import (
	"fmt"
	"math/rand"
	"sf_graphs/bintree"
	"sf_graphs/graph"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//Создаем бинарное дерево и заполняем его значениями.
	tree := bintree.NewTree()
	for j := 0; j < 50; j++ {
		tree.AddNode(rand.Intn(20) - 10)
	}

	//Печатаем элементы дерева.
	tree.PrintTree()
	fmt.Printf("\n")

	//Находим узлы со значением 2.
	nodes, err := tree.FindAll(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nНайденные узлы:\n")
	for i := range nodes {
		fmt.Println(nodes[i])
	}

	//Удаляем все узлы со значением 2.
	tree.DelAll(2)

	//Проверяем, что нужные узлы удалились.
	nodes, err = tree.FindAll(2)
	if err != nil {
		fmt.Println(err)
	}
	for i := range nodes {
		fmt.Println(nodes[i])
	}

	//Составляем структуру автомобильных дорог Приморского края
	//в виде неориентированного взвешенного графа.
	road := graph.NewGraph()

	vdk := road.NewNode("Владивосток")
	uss := road.NewNode("Уссурийск")
	pgr := road.NewNode("Пограничный")
	hrl := road.NewNode("Хороль")
	kmr := road.NewNode("Камень-Рыболов")
	ars := road.NewNode("Арсеньев")
	mal := road.NewNode("Малиново")
	che := road.NewNode("Черниговка")
	spd := road.NewNode("Спасск-Дальний")
	lsz := road.NewNode("Лесозаводск")
	dlr := road.NewNode("Дальнереченск")

	road.NewEdgeUndir(vdk, uss, 98)
	road.NewEdgeUndir(uss, pgr, 96)
	road.NewEdgeUndir(uss, hrl, 74)
	road.NewEdgeUndir(uss, che, 91)
	road.NewEdgeUndir(uss, ars, 145)
	road.NewEdgeUndir(pgr, kmr, 71)
	road.NewEdgeUndir(hrl, kmr, 36)
	road.NewEdgeUndir(hrl, che, 51)
	road.NewEdgeUndir(ars, che, 82)
	road.NewEdgeUndir(ars, spd, 112)
	road.NewEdgeUndir(ars, mal, 231)
	road.NewEdgeUndir(mal, dlr, 92)
	road.NewEdgeUndir(che, spd, 40)
	road.NewEdgeUndir(spd, lsz, 133)
	road.NewEdgeUndir(lsz, dlr, 70)

	//Выводим структуру в консоль.
	for i := range road.Nodes {
		fmt.Printf("\n  %v   %v", road.Nodes[i].Value, road.Nodes[i].Matrix)
	}

	//Ищем Лесозаводск.
	//Поиск вширину без учета весов.
	n, err := road.BFS("Лесозаводск")
	if err != nil {
		fmt.Printf("\n\n%v", err)
		return
	}
	fmt.Printf("\n\nНайден узел: %v", n.Value)

	//Ищем кратчайший путь между Владивостоком и Дальнереченском.
	//Поиск по алгоритму Дейкстры с учетом весов.
	way, err := road.Dijkstra(vdk, dlr)
	if err != nil {
		fmt.Printf("\n\n%v", err)
		return
	}
	fmt.Printf("\n\n Путь между узлами:")
	for i := range way {
		fmt.Printf("\n  %v", way[i].Value)
	}

	//Составляем структуру из имен каких-то людей
	//в виде ориентированного взвешенного графа.
	peoples := graph.NewGraph()

	vasya := peoples.NewNode("Вася")
	katya := peoples.NewNode("Катя")
	olya := peoples.NewNode("Оля")
	petya := peoples.NewNode("Петя")
	sasha := peoples.NewNode("Саша")
	nadya := peoples.NewNode("Надя")
	kostya := peoples.NewNode("Костя")

	peoples.NewEdge(vasya, katya, 33)
	peoples.NewEdge(vasya, olya, 4)
	peoples.NewEdge(olya, vasya, 8)
	peoples.NewEdge(vasya, petya, 50)
	peoples.NewEdge(katya, olya, 20)
	peoples.NewEdge(petya, katya, 4)
	peoples.NewEdge(petya, sasha, 13)
	peoples.NewEdge(sasha, petya, 14)
	peoples.NewEdge(nadya, petya, 12)
	peoples.NewEdge(nadya, sasha, 22)
	peoples.NewEdge(sasha, nadya, 1)
	peoples.NewEdge(kostya, petya, 3)

	//Ищем кратчайший путь от Васи к Наде.
	way, err = peoples.Dijkstra(vasya, nadya)
	if err != nil {
		fmt.Printf("\n\n%v", err)
		return
	}
	fmt.Printf("\n\n Путь между узлами:")
	for i := range way {
		fmt.Printf("\n  %v", way[i].Value)
	}

	//И путь обратно от Нади к Васе.
	way, err = peoples.Dijkstra(nadya, vasya)
	if err != nil {
		fmt.Printf("\n\n%v", err)
		return
	}
	fmt.Printf("\n\n Путь между узлами:")
	for i := range way {
		fmt.Printf("\n  %v", way[i].Value)
	}

	//Пытаемся найти несуществующий путь.
	way, err = peoples.Dijkstra(nadya, kostya)
	if err != nil {
		fmt.Printf("\n\n%v", err)
		return
	}
	fmt.Printf("\n\n Путь между узлами:")
	for i := range way {
		fmt.Printf("\n  %v", way[i].Value)
	}
}
