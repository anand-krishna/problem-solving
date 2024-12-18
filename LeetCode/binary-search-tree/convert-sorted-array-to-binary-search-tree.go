/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	midIndex := len(nums) / 2
	t := &TreeNode{
		Val:   nums[midIndex],
		Left:  nil,
		Right: nil,
	}
	t.Left = insertIntoBST(nums[:midIndex])
	t.Right = insertIntoBST(nums[midIndex+1:])
	return t
}

func sortedArrayToBST(nums []int) *TreeNode {
	// Sorted Array := Inorder representation of the BST.
	return insertIntoBST(nums)
}