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

func (l list) Equals(l2 list) bool {
	if len(l) != len(l2) {
		return false
	}
	for i := range l {
		if !equal(l[i], l2[i]) {
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
