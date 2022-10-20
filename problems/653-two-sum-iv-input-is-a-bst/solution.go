/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	memo := map[int]bool{}

	return findTargetRecursively(root, k, memo)
}

func findTargetRecursively(node *TreeNode, k int, memo map[int]bool) bool {
	if node == nil {
		return false
	}

	if _, ok := memo[k-node.Val]; ok {
		return true
	} else {
		memo[node.Val] = true
	}

	return findTargetRecursively(node.Left, k, memo) || findTargetRecursively(node.Right, k, memo)
}
