package models

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack:=[]*TreeNode{}
	res := []int{}
	p := root
	for len(stack)>0 || p!=nil {
		for p!=nil {
			stack = append(stack, p)
			res = append(res, p.Val)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p = p.Right
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack:=[]*TreeNode{}
	res := []int{}
	p := root
	var pre *TreeNode
	for p!=nil {
		stack = append(stack, p)
		p = p.Left
	}
	for len(stack)>0 {
		p = stack[len(stack)-1]
		if p.Right==nil || p.Right==pre {
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			pre = p
		} else {
			p = p.Right
			for p!=nil {
				stack = append(stack, p)
				p = p.Left
			}
		}
	}
	return res
}