package goxtream

import (
	"strconv"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
)

type Category struct {
	Id       int
	Name     string
	Parent   int
	Children []*Category
}

func parseCategory(src dtos.CategoryResult) (*Category, error) {
	id, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, err
	}

	return &Category{
		Id:       id,
		Name:     src.CategoryName,
		Parent:   src.ParentId,
		Children: make([]*Category, 0),
	}, nil
}

func parseCategories(srcs []dtos.CategoryResult) ([]*Category, error) {
	categories := make([]*Category, 0)
	for _, c := range srcs {
		category, err := parseCategory(c)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	for _, category := range categories {
		if category.Parent == 0 {
			continue
		}
		for _, parent := range categories {
			parent.Children = append(parent.Children, category)
		}
	}

	filtered := categories[:0]
	for _, category := range categories {
		if category.Parent == 0 {
			filtered = append(filtered, category)
		}
	}

	return filtered, nil
}
