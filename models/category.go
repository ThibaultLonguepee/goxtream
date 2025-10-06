package models

type Category struct {
	Id       int
	Name     string
	Parent   int
	Children []*Category
}
