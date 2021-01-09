import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	sep   string
	null  string
	items []string
}

func Constructor() Codec {
	return Codec{
		sep:  ",",
		null: "null",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.serializeHelper(root)
	res := strings.Join(this.items, this.sep)
	this.items = this.items[:0]
	return res
}

func (this *Codec) serializeHelper(root *TreeNode) {
	if root == nil {
		this.items = append(this.items, this.null)
		return
	}
	this.items = append(this.items, strconv.Itoa(root.Val))
	this.serializeHelper(root.Left)
	this.serializeHelper(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.items = strings.Split(data, this.sep)
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if len(this.items) == 0 {
		return nil
	}
	item := this.items[0]
	this.items = this.items[1:]
	if item == this.null {
		return nil
	}
	val, err := strconv.Atoi(item)
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: val}
	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */