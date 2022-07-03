package collections

import "reflect"

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
