package bintree

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BinaryTree struct {
	Root *Node
}

//Ошибка при передаче значения, отсутствующего в дереве
var ErrorNumberIsMissing = fmt.Errorf("Число отсутствует в коллекции")

//Создать узел
func (bt *BinaryTree) NewNode(parent *Node, val int) *Node {
	return &Node{val, nil, nil, parent}
}

//Создать дерево
func NewTree() *BinaryTree {
	return &BinaryTree{nil}
}

//Добавить новый узел
func (bt *BinaryTree) AddNode(val int) {
	current := bt.Root

	if bt.Root == nil {
		bt.Root = &Node{val, nil, nil, nil}
		return
	}

	for {
		if val <= current.Value {
			if current.Left == nil {
				newNode := bt.NewNode(current, val)
				current.Left = newNode
				return
			}
			current = current.Left
		}
		if val > current.Value {
			if current.Right == nil {
				newNode := bt.NewNode(current, val)
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

//Вывести в консоль все элементы дерева
func (bt BinaryTree) PrintTree() {
	if bt.Root == nil {
		return
	}

	fmt.Printf("%d\t", bt.Root.Value)

	leftTree := BinaryTree{bt.Root.Left}
	righrTree := BinaryTree{bt.Root.Right}

	leftTree.PrintTree()
	righrTree.PrintTree()
}

//Найти ближайший к корню узел по значению
func (bt BinaryTree) FindNode(val int) (*Node, error) {
	if bt.Root == nil {
		return nil, ErrorNumberIsMissing
	}

	leftTree := BinaryTree{bt.Root.Left}
	rightTree := BinaryTree{bt.Root.Right}

	if val < bt.Root.Value {
		return leftTree.FindNode(val)
	}
	if val > bt.Root.Value {
		return rightTree.FindNode(val)
	}
	return bt.Root, nil
}

//Найти все повторяющиеся узлы по значению
func (bt BinaryTree) FindAll(val int) (nodes []*Node, err error) {
	n := bt.Root
	for {
		tree := BinaryTree{n}
		n, err = tree.FindNode(val)
		if err != nil {
			break
		}
		nodes = append(nodes, n)
		n = n.Left
	}
	if len(nodes) == 0 {
		return nil, ErrorNumberIsMissing
	}
	return nodes, nil
}

//Удалить ближайший к корню узел по значению
func (bt *BinaryTree) DelNode(val int) error {
	n, err := bt.FindNode(val)
	if err != nil {
		return err
	}
	//левый есть, правого нет
	if n.Left != nil {
		if n.Right == nil {
			n.Left.Parent = n.Parent
			if n != bt.Root {
				if n == n.Parent.Left {
					n.Parent.Left = n.Left
					return nil
				}
				n.Parent.Right = n.Left
				return nil
			}
			bt.Root = n.Left
			return nil
		}
		//левый есть, правый есть
		heir := n.Right
		if heir.Left != nil {
			for {
				heir = heir.Left
				if heir.Left == nil {
					break
				}
			}
			if heir.Right != nil {
				heir.Right.Parent = heir.Parent
			}
			heir.Parent.Left = heir.Right
			n.Value = heir.Value
			heir = nil
			return nil
		}
		n.Right = heir.Right
		heir.Right.Parent = n
		n.Value = heir.Value
		heir = nil
		return nil
	}
	//левого нет, правый есть
	if n.Right != nil {
		n.Right.Parent = n.Parent
		if n != bt.Root {
			if n == n.Parent.Left {
				n.Parent.Left = n.Right
				return nil
			}
			n.Parent.Right = n.Right
			return nil
		}
		bt.Root = n.Right
		return nil
	}
	//левого нет, правого нет
	if n == bt.Root {
		bt.Root = nil
		return nil
	}
	if n == n.Parent.Left {
		n.Parent.Left = nil
		return nil
	}
	n.Parent.Right = nil
	return nil
}

//Удалить все повторяющиеся узлы по значению
func (bt *BinaryTree) DelAll(val int) {
	for {
		err := bt.DelNode(val)
		if err != nil {
			break
		}
	}
}
