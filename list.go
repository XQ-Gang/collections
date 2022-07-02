package collections

import (
	"reflect"
)

type list []any

func List(elems ...any) list {
	l := list(elems)
	return l
}

func (l *list) Append(elem any) {
	*l = append(*l, elem)
}

func (l *list) Pop() any {
	size := len(*l)
	if size == 0 {
		panic("pop from empty list")
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
	case list:
		return elem1.(list).Equals(elem2.(list))
	}
	return elem1 == elem2
}

func (l list) Equals(_l list) bool {
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

func (l list) Index(elem any) any {
	for i, item := range l {
		if equal(item, elem) {
			return i
		}
	}
	return -1
}

func (l *list) Extend(_l list) {
	*l = append(*l, _l...)
}

func (l list) Count(elem any) any {
	count := 0
	for _, item := range l {
		if equal(item, elem) {
			count++
		}
	}
	return count
}
