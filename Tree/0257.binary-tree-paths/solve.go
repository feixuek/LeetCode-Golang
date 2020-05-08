import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	head := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{head}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	res := make([]string, 0)
	for _, v := range left {
		res = append(res, head+"->"+v)
	}
	for _, v := range right {
		res = append(res, head+"->"+v)
	}
	return res
}