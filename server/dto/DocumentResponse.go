package dto

import (
	"github.com/google/uuid"
)

type DocumentResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"documentName"`
	Slug        string    `json:"slug"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Font        string    `json:"font"`
}
