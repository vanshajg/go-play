package models

type Comment struct {
	by      string
	id      int
	parent  int
	text    string
	time    int
	type_of string
}
