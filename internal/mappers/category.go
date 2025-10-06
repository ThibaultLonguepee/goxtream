package mappers

import (
	"strconv"

	"github.com/thibaultlonguepee/goxtream/internal/dtos"
	"github.com/thibaultlonguepee/goxtream/models"
)

func ParseCategory(src dtos.CategoryResult) (*models.Category, error) {
	id, err := strconv.Atoi(src.CategoryId)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:       id,
		Name:     src.CategoryName,
		Parent:   src.ParentId,
		Children: make([]*models.Category, 0),
	}, nil
}

func ParseCategories(srcs []dtos.CategoryResult) ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	for _, c := range srcs {
		category, err := ParseCategory(c)
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
