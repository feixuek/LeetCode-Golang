func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := []int{}
	arr2 := []int{}
	inorder(root1, &arr1)
	inorder(root2, &arr2)
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {
	idx1, idx2 := 0, 0
	res := []int{}
	for idx1 < len(arr1) && idx2 < len(arr2) {
		if arr1[idx1] <= arr2[idx2] {
			res = append(res, arr1[idx1])
			idx1++
		} else {
			res = append(res, arr2[idx2])
			idx2++
		}
	}
	if idx1 < len(arr1) {
		res = append(res, arr1[idx1:]...)
	}
	if idx2 < len(arr2) {
		res = append(res, arr2[idx2:]...)
	}
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	(*res) = append(*res, root.Val)
	inorder(root.Right, res)
}