package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumOfLeftLeaves(&TreeNode{1, nil, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}

	stack := []*TreeNode{root}
	i := 1
	for i < len(nums) && len(stack) > 0 {
		node := stack[0]
		node.Left = &TreeNode{nums[i], nil, nil}
		stack = append(stack, node.Left)
		i++
		if i >= len(nums) {
			break
		}
		node.Right = &TreeNode{nums[i], nil, nil}
		stack = append(stack, node.Right)
		i++
		if i >= len(nums){
			break
		}
		stack = stack[1:]
	}

	return root
}
