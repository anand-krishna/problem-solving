/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	if root == nil {
		return sum
	}

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	sum += rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)

	return sum
}