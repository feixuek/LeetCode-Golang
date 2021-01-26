type Node struct {
	Key  int
	Val  int
	Next *Node
	Pre  *Node
}

type LinkedList struct {
	Size int
	Cap  int
	Head *Node
	Tail *Node
}

func (list *LinkedList) Remove(elem *Node) {
	elem.Pre.Next = elem.Next
	elem.Next.Pre = elem.Pre
	list.Size--
}

func (list *LinkedList) DeletFisrt() *Node {
	if list.Head.Next == nil {
		return nil
	}
	elem := list.Head.Next
	list.Remove(elem)
	return elem
}

func (list *LinkedList) AddLast(elem *Node) {
	elem.Pre = list.Tail.Pre
	list.Tail.Pre.Next = elem
	elem.Next = list.Tail
	list.Tail.Pre = elem
	list.Size++
}

type LRUCache struct {
	List   *LinkedList
	Filter map[int]*Node
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		List: &LinkedList{
			Size: 0,
			Cap:  capacity,
			Head: &Node{},
			Tail: &Node{},
		},
		Filter: make(map[int]*Node),
	}
	res.List.Head.Next = res.List.Tail
	res.List.Tail.Pre = res.List.Head
	return res
}

func (this *LRUCache) Get(key int) int {
	if this == nil {
		return -1
	}
	elem, ok := this.Filter[key]
	if !ok {
		return -1
	}
	this.List.Remove(elem)
	this.List.AddLast(elem)
	return elem.Val
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{
		Key:  key,
		Val:  value,
		Next: new(Node),
		Pre:  new(Node),
	}
	if elem, ok := this.Filter[key]; ok {
		this.List.Remove(elem)
		delete(this.Filter, key)
	}
	if this.List.Size >= this.List.Cap {
		elem := this.List.DeletFisrt()
		if elem != nil {
			delete(this.Filter, elem.Key)
		}
	}
	this.List.AddLast(node)
	this.Filter[key] = node
}
