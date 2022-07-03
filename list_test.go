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

	list = List(1, List(1, List(1, List(1))))
	got = fmt.Sprintf("%v", list)
	want = "[1 [1 [1 [1]]]]"
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

func TestList_Equals(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2}, List(1, List()))
	got := list.Equals(list)
	want := true
	if got != want {
		t.Errorf("list.Equals(list) = %v, want %v", got, want)
	}

	got = list.Equals(List())
	want = false
	if got != want {
		t.Errorf("list.Equals(List()) = %v, want %v", got, want)
	}
}

func TestList_Index(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2}, List(1))
	got := list.Index([2]int{1, 2})
	want := 3
	if got != want {
		t.Errorf("list.Index([2]int{1, 2}) = %v, want %v", got, want)
	}

	got = list.Index(5)
	want = -1
	if got != want {
		t.Errorf("list.Index(5) = %v, want %v", got, want)
	}

	got = list.Index(List(1))
	want = 4
	if got != want {
		t.Errorf("list.Index(List(1)) = %v, want %v", got, want)
	}
}

func TestList_Extend(t *testing.T) {
	list := List(1)
	list.Extend(List(2))
	want := List(1, 2)
	if !list.Equals(want) {
		t.Errorf("list.Extend(List(2)), list = %v, want %v", list, want)
	}

	list.Extend(List(List(3)))
	want = List(1, 2, List(3))
	if !list.Equals(want) {
		t.Errorf("list.Extend(List(List(3))), list = %v, want %v", list, want)
	}
}

func TestList_Count(t *testing.T) {
	list := List(1, 1, 1, 1)
	got := list.Count(1)
	want := 4
	if got != want {
		t.Errorf("list.Count(1) = %v, want %v", got, want)
	}

	got = list.Count(0)
	want = 0
	if got != want {
		t.Errorf("list.Count(0) = %v, want %v", got, want)
	}

	list = List(1, List(1, List(1)), List(1, List(1)))
	got = list.Count(List(1, List(1)))
	want = 2
	if got != want {
		t.Errorf("list.Count(List(1, List(1))) = %v, want %v", got, want)
	}
}

func TestList_Clear(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2}, List(1))
	list.Clear()
	want := List()
	if !list.Equals(want) {
		t.Errorf("list.Clear(), list = %v, want %v", list, want)
	}
}

func TestList_Copy(t *testing.T) {
	list := List(1, 2, 3, 4)
	copied := list.Copy()
	list[1] = 0

	got := list[1]
	want := 0
	if got != want {
		t.Errorf("list.Copy(), list[1] = %v, want %v", copied[1], want)
	}

	got = copied[1]
	want = 2
	if got != want {
		t.Errorf("list.Copy(), copied[1] = %v, want %v", copied[1], want)
	}
}

func TestList_Reverse(t *testing.T) {
	list := List(1, 2, 3, 4)
	list.Reverse()
	want := List(4, 3, 2, 1)
	if !list.Equals(want) {
		t.Errorf("list.Reverse(), list = %v, want %v", list, want)
	}
}

func TestList_Insert(t *testing.T) {
	list := List(1)
	list.Insert(0, "a")
	list.Insert(1, "b")
	list.Insert(3, "c")
	list.Insert(-10, "d")
	list.Insert(10, "e")
	want := List("d", "a", "b", 1, "c", "e")
	if !list.Equals(want) {
		t.Errorf("list.Insert(), list = %v, want %v", list, want)
	}
}

func TestList_Remove(t *testing.T) {
	list := List(1, 2, 3, 4, 5, 4, 3, 2, 1)
	list.Remove(1)
	list.Remove(5)
	list.Remove(1)
	want := List(2, 3, 4, 4, 3, 2)
	if !list.Equals(want) {
		t.Errorf("list.Remove(), list = %v, want %v", list, want)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("list.Remove(0) = nil, want panic(\"List.Remove(x): x not in List\")")
		}
	}()
	list.Remove(0)
}
