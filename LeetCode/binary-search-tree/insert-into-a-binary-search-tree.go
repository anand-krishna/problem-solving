/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	insert(root, val)
	return root
}

func insert(root *TreeNode, val int) {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return
		}
		insertIntoBST(root.Right, val)
	} else if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return
		}
		insertIntoBST(root.Left, val)
	}
}