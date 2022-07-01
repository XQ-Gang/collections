package collections

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
