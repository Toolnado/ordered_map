package ordered_map

type OrderedMap struct {
	root *Node
	size int
}
type Node struct {
	Key        int
	Value      int
	LeftChild  *Node
	RigthChild *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:        key,
		Value:      value,
		LeftChild:  nil,
		RigthChild: nil,
	}
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	m.size++

	if m.root == nil {
		m.root = NewNode(key, value)
		return
	}

	node := m.root
	for {
		switch {
		case node.Key == key:
			node.Value = value
			return

		case node.Key > key:
			if node.LeftChild == nil {
				node.LeftChild = NewNode(key, value)
				return
			} else {
				node = node.LeftChild
			}

		case node.Key < key:
			if node.RigthChild == nil {
				node.RigthChild = NewNode(key, value)
				return
			} else {
				node = node.RigthChild
			}
		}
	}
}

func (m *OrderedMap) Erase(key int) {
	var (
		x = m.root
		y *Node
	)

	for x != nil {
		if x.Key == key {
			break
		} else {
			y = x
			if x.Key > key {
				x = x.LeftChild
			} else {
				x = x.RigthChild
			}
		}
	}

	if x == nil {
		return
	}

	if x.RigthChild == nil {
		if y == nil {
			m.root = x.LeftChild
		} else {
			if y.Key > x.Key {
				y.LeftChild = x.LeftChild
			} else {
				y.RigthChild = x.LeftChild
			}
		}
	} else {
		left := x.RigthChild
		y = nil

		for left.LeftChild != nil {
			y = left
			left = left.LeftChild
		}

		if y != nil {
			y.LeftChild = left.RigthChild
		} else {
			x.RigthChild = left.RigthChild
		}

		x.Key = left.Key
		x.Value = left.Value
	}

	m.size--
}

func (m *OrderedMap) Contains(key int) bool {
	node := m.root
	for node != nil {
		switch {
		case node.Key == key:
			return true

		case node.Key > key:
			node = node.LeftChild

		case node.Key < key:
			node = node.RigthChild
		}
	}
	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.use(m.root, action)
}

func (m *OrderedMap) use(node *Node, action func(int, int)) {
	if node.LeftChild != nil {
		m.use(node.LeftChild, action)
	}
	action(node.Key, node.Value)
	if node.RigthChild != nil {
		m.use(node.RigthChild, action)
	}
}
