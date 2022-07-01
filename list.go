package collections

type list []any

func List(elems ...any) list {
	l := list(elems)
	return l
}
