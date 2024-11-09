package ordered_map

type OrderedMap struct {
	// need to implement
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{} // need to implement
}

func (m *OrderedMap) Insert(key, value int) {
	// need to implement
}

func (m *OrderedMap) Erase(key int) {
	// need to implement
}

func (m *OrderedMap) Contains(key int) bool {
	return false // need to implement
}

func (m *OrderedMap) Size() int {
	return 0 // need to implement
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	// need to implement
}
