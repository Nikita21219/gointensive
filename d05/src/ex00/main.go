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

func countOne(node *TreeNode) int {
	count := 0
	return *countOneByNode(node, &count)
}

func countOneByNode(root *TreeNode, count *int) *int {
	if root == nil {
		return count
	}
	countOneByNode(root.Left, count)
	if root.HasToy {
		*count++
	}
	//fmt.Println("Has Toy:", root.HasToy)
	countOneByNode(root.Right, count)
	return count
}

func areToysBalanced(root *TreeNode) bool {
	nodes := []*TreeNode{root.Left, root.Right}
	resultCh := make(chan int, 2)

	for _, node := range nodes {
		node := node
		go func() {
			resultCh <- countOne(node)
		}()
	}

	var pair []int
	for range nodes {
		pair = append(pair, <-resultCh)
	}
	close(resultCh)
	return pair[0] == pair[1]
}

func main() {
	fmt.Printf("1 example (true): ")
	{
		// 1 example
		root := createNode(false)
		n2 := createNode(false)
		n3 := createNode(true)
		n4 := createNode(false)
		n5 := createNode(true)

		root.Left = n2
		root.Right = n3
		n2.Left = n4
		n2.Right = n5
		fmt.Println(areToysBalanced(root))
	}

	fmt.Printf("2 example (true): ")
	{
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
		fmt.Println(areToysBalanced(root))
	}

	fmt.Printf("3 example (false): ")
	{
		root := createNode(true)
		n2 := createNode(true)
		n3 := createNode(false)

		root.Left = n2
		root.Right = n3
		fmt.Println(areToysBalanced(root))
	}

	fmt.Printf("4 example (false): ")
	{
		root := createNode(false)
		n2 := createNode(true)
		n3 := createNode(false)
		n4 := createNode(true)
		n5 := createNode(true)

		root.Left = n2
		root.Right = n3
		n2.Right = n4
		n3.Right = n5
		fmt.Println(areToysBalanced(root))
	}
}
