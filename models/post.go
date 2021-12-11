package models

type Post struct {
	by         string
	decendents int
	id         int
	kids       []int
}
