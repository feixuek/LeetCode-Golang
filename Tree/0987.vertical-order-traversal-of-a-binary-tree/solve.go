import "sort"

type priorityValue struct {
	Val   int
	Level int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	filter := make(map[int][]priorityValue)
	helper(root, &filter, 0, 0)
	keys := []int{}
	for k := range filter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := make([][]int, len(keys))
	for i := 0; i < len(keys); i++ {
		pv := filter[keys[i]]
		sort.Slice(pv, func(i, j int) bool {
			if pv[i].Level == pv[j].Level {
				return pv[i].Val < pv[j].Val
			}
			return pv[i].Level < pv[j].Level
		})
		for j := 0; j < len(pv); j++ {
			res[i] = append(res[i], pv[j].Val)
		}
	}
	return res
}

func helper(root *TreeNode, res *map[int][]priorityValue, coordinate, level int) {
	if root == nil {
		return
	}
	(*res)[coordinate] = append((*res)[coordinate], priorityValue{Val: root.Val, Level: level})
	level++
	helper(root.Left, res, coordinate-1, level)
	helper(root.Right, res, coordinate+1, level)
}