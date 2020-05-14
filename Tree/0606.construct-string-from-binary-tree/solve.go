import "strconv"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	res := strconv.Itoa(t.Val)
	left := tree2str(t.Left)
	right := tree2str(t.Right)
	switch {
	case left == "" && right == "":
		return res
	case left == "":
		return res + "()" + "(" + right + ")"
	case right == "":
		return res + "(" + left + ")"
	}
	return res + "(" + left + ")(" + right + ")"
}