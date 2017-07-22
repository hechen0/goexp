package main

import "fmt"

func main() {
	root := &TreeNode{1,
					  &TreeNode{2, nil, nil},
					  &TreeNode{3, nil, nil},
	}

	postOrder(root)
}

func inOrder(head *TreeNode) {
	s := make([]*TreeNode, 0)

	cur := head

	for len(s) != 0 || cur != nil {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}

		if len(s) != 0 {
			cur = s[len(s)-1]
			fmt.Println(cur.Val)
			s = s[:len(s)-1]
			cur = cur.Right
		}
	}
}

func preOrder(head *TreeNode) {
	s := make([]*TreeNode, 0)
	cur := head
	for len(s) != 0 || cur != nil {
		for cur != nil {
			fmt.Println(cur.Val)
			s = append(s, cur)
			cur = cur.Left
		}

		if len(s) != 0 {
			cur = s[len(s)-1]
			s = s[:len(s)-1]
			cur = cur.Right
		}
	}
}

func postOrder(head *TreeNode) {
	s := make([]*TreeNode, 0)
	s = append(s, head)

	var pre, cur *TreeNode

	for len(s) != 0 {
		cur = s[len(s)-1]

		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			fmt.Println(cur.Val)
			s = s[:len(s)-1]
			pre = cur
		} else {
			if cur.Right != nil {
				s = append(s, cur.Right)
			}

			if cur.Left != nil {
				s = append(s, cur.Left)
			}
		}
	}
}
