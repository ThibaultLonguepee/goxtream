package dtos

type CategoryResult struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentId     int    `json:"parent_id"`
}
