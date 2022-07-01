package collections

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	if list[0] != 1 {
		t.Errorf("list[0] = %v, want %v", list[0], 1)
	}
	if list[1] != 3.14 {
		t.Errorf("list[1] = %v, want %v", list[1], 3.14)
	}
	if list[2] != "2" {
		t.Errorf("list[2] = %v, want %v", list[2], "2")
	}
	if list[3] != [2]int{1, 2} {
		t.Errorf("list[3] = %v, want %v", list[3], [2]int{1, 2})
	}
}

func TestList_Size(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	got := len(list)
	want := 4
	if got != want {
		t.Errorf("len(list) = %v, want %v", got, want)
	}
}

func TestList_String(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	got := fmt.Sprintf("%v", list)
	want := "[1 3.14 2 [1 2]]"
	if got != want {
		t.Errorf("%%v = %v, want %v", got, want)
	}
}

func TestList_Append(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	list.Append(5)
	got := list[len(list)-1]
	want := 5
	if got != want {
		t.Errorf("list.Append(5), list[-1] = %v, want %v", got, want)
	}
}

func TestList_Pop(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	got := list.Pop()
	want := [2]int{1, 2}
	if got != want {
		t.Errorf("list.Pop() = %v, want %v", got, want)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("elist.Pop() = nil, want panic(\"pop from empty list\")")
		}
	}()
	elist := List()
	elist.Pop()
}
