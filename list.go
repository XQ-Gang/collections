package collections

import (
	"reflect"
)

type _list []any

func List(elems ...any) _list {
	l := _list(elems)
	return l
}

func (l *_list) Append(elem any) {
	*l = append(*l, elem)
}

func (l *_list) Pop() any {
	size := len(*l)
	if size == 0 {
		panic("pop from empty _list")
	}
	last := (*l)[size-1]
	*l = (*l)[:size-1]
	return last
}

func equal(elem1, elem2 any) bool {
	if reflect.TypeOf(elem1) != reflect.TypeOf(elem2) {
		return false
	}
	switch elem1.(type) {
	case _list:
		return elem1.(_list).Equals(elem2.(_list))
	}
	return elem1 == elem2
}

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

func (l _list) Index(elem any) any {
	for i, item := range l {
		if equal(item, elem) {
			return i
		}
	}
	return -1
}

func (l *_list) Extend(_l _list) {
	*l = append(*l, _l...)
}

func (l _list) Count(elem any) any {
	count := 0
	for _, item := range l {
		if equal(item, elem) {
			count++
		}
	}
	return count
}

func (l *_list) Clear() {
	*l = List()
}

func (l _list) Copy() _list {
	_l := make(_list, len(l))
	copy(_l, l)
	return _l
}
