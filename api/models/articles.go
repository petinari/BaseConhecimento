package models

import "encoding/json"

type Articles struct {
	ID          uint64          `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	ImageUrl    string          `json:"imageUrl,omitempty"`
	Content     json.RawMessage `json:"content,omitempty"`
	UserId      uint64          `json:"userId,omitempty"`
	CategorieId uint64          `json:"categorieId,omitempty"`
}
