package collections

import (
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

func TestListSize(t *testing.T) {
	list := List(1, 3.14, "2", [...]int{1, 2})
	got := len(list)
	want := 4
	if got != want {
		t.Errorf("len(list) = %v, want %v", got, want)
	}
}
