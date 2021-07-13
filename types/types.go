package types

// ToDo struct
type ToDo struct {
	ID      int
	Index   int
	Caption string
	Excute  bool
}

// LoadAllToDoData for template data
type LoadAllToDoData struct {
	ToDos []*ToDo
}
