// Package handler contains all item business logic
package handler

import (
	"PaginationPlayground/internal/models"
	"PaginationPlayground/internal/persist"
)

type ItemHandler struct {
	items persist.ItemRepository
}

func NewItemHandler(repo persist.ItemRepository) *ItemHandler {
	return &ItemHandler{repo}
}

func (h *ItemHandler) SearchForItems(itemName string) ([]models.SearchItem, error) {
	items, err := h.items.GetItem(itemName)
	if err != nil {
		return nil, err
	}
	return items, nil
}
