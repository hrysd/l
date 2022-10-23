/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	answer := []string{}

	getPaths(root, "", &answer)

	return answer
}

func getPaths(node *TreeNode, s string, list *[]string) {
	v := strconv.Itoa(node.Val)
	s = s + v

	if node.Left == nil && node.Right == nil {
		*list = append(*list, s)
		return
	}

	if node.Left != nil {
		getPaths(node.Left, s+"->", list)
	}
	if node.Right != nil {
		getPaths(node.Right, s+"->", list)
	}
}


