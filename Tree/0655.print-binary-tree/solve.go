import (
	"math"
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	h := height(root)
	column := int(math.Pow(2.0, float64(h))) - 1
	res := make([][]string, h)
	for i := 0; i < h; i++ {
		res[i] = append(res[i], make([]string, column)...)
	}
	helper(root, 0, 0, column-1, res)
	return res
}

func helper(root *TreeNode, row, left, right int, res [][]string) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	res[row][mid] = strconv.Itoa(root.Val)
	helper(root.Left, row+1, left, mid-1, res)
	helper(root.Right, row+1, mid+1, right, res)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}