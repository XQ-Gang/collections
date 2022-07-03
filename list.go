package collections

type _list []any

func List(elems ...any) _list {
	l := _list(elems)
	return l
}

// Append object to the end of the list.
func (l *_list) Append(elem any) {
	*l = append(*l, elem)
}

// Clear Remove all items from list.
func (l *_list) Clear() {
	*l = List()
}

// Copy Return a shallow copy of the list.
func (l _list) Copy() _list {
	_l := make(_list, len(l))
	copy(_l, l)
	return _l
}

// Count Return number of occurrences of value.
func (l _list) Count(elem any) any {
	count := 0
	for _, item := range l {
		if equal(item, elem) {
			count++
		}
	}
	return count
}

// Equals Return whether two lists are equal
func (l _list) Equals(_l _list) bool {
	if len(l) != len(_l) {
		return false
	}
	for i := range l {
		if !equal(l[i], _l[i]) {
			return false
		}
	}
	return true
}

// Extend list by appending elements from the iterable.
func (l *_list) Extend(_l _list) {
	*l = append(*l, _l...)
}

// Index Return first index of value.
func (l _list) Index(elem any) any {
	for i, item := range l {
		if equal(item, elem) {
			return i
		}
	}
	return -1
}

// Insert object before index.
func (l *_list) Insert(index int, elem any) {
	size := len(*l)
	if index < 0 {
		index = size + index
		if index < 0 {
			index = 0
		}
	}
	if index > size {
		index = size
	}
	*l = append((*l)[:index], append(List(elem), (*l)[index:]...)...)
}

// Pop Remove and return item at index (default last).
func (l *_list) Pop() any {
	size := len(*l)
	if size == 0 {
		panic("pop from empty List")
	}
	last := (*l)[size-1]
	*l = (*l)[:size-1]
	return last
}

// Remove first occurrence of value.
func (l *_list) Remove(elem any) {
	for i, item := range *l {
		if equal(item, elem) {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
	panic("List.Remove(x): x not in List")
}

// Reverse *IN PLACE*.
func (l *_list) Reverse() {
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}
