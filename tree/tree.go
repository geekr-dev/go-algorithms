package tree

type Tree interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}
