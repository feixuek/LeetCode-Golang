import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	counter := make(map[string]int)
	res := []*TreeNode{}
	helper(root, &counter, &res)
	return res
}

func helper(root *TreeNode, counter *map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	left := helper(root.Left, counter, res)
	right := helper(root.Right, counter, res)
	key := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
	if c, ok := (*counter)[key]; ok && c == 1 {
		*res = append(*res, root)
	}
	(*counter)[key]++
	return key
}