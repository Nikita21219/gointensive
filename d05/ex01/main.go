package main

import (
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func createNode(value bool) *TreeNode {
	node := new(TreeNode)
	node.HasToy = value
	node.Right = nil
	node.Left = nil
	return node
}

func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return nil
	}

	var result []bool
	queue := []*TreeNode{root}
	level := 1

	for len(queue) > 0 {
		size := len(queue)
		levelValues := make([]bool, size)

		for i := 0; i < size; i++ {
			node := queue[i]

			if level%2 == 0 {
				levelValues[i] = node.HasToy
			} else {
				levelValues[size-i-1] = node.HasToy
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValues...)
		queue = queue[size:]
		level++
	}

	return result
}

func main() {
	{
		fmt.Println("1 example")
		root := createNode(true)
		n2 := createNode(true)
		n3 := createNode(false)
		n4 := createNode(true)
		n5 := createNode(false)
		n6 := createNode(true)
		n7 := createNode(true)

		root.Left = n2
		root.Right = n3
		n2.Left = n4
		n2.Right = n5
		n3.Left = n6
		n3.Right = n7
		rightRes := []bool{true, true, false, true, true, false, true}
		fmt.Println("Right result: ", rightRes)
		fmt.Println("Result:       ", unrollGarland(root))
	}
	fmt.Println()
	{
		fmt.Println("2 example")
		root := createNode(true)
		n2 := createNode(true)
		n3 := createNode(false)
		n4 := createNode(true)
		n5 := createNode(false)
		n6 := createNode(true)
		n7 := createNode(true)
		n8 := createNode(true)
		n9 := createNode(false)
		n10 := createNode(true)
		n11 := createNode(true)
		n12 := createNode(true)
		n13 := createNode(false)
		n14 := createNode(true)
		n15 := createNode(true)

		root.Left = n2
		root.Right = n3
		n2.Left = n4
		n2.Right = n5
		n3.Left = n6
		n3.Right = n7
		n4.Left = n8
		n4.Right = n9
		n5.Left = n10
		n5.Right = n11
		n6.Left = n12
		n6.Right = n13
		n7.Left = n14
		n7.Right = n15

		rightRes := []bool{true, true, false, true, true, false, true, true, false, true, true, true, false, true, true}
		fmt.Println("Right result: ", rightRes)
		fmt.Println("Result:       ", unrollGarland(root))
	}
}
