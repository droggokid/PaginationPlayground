// Package service contains item related business logic
package service

import (
	"PaginationPlayground/internal/client"
	"PaginationPlayground/internal/models"
	"PaginationPlayground/internal/persist"
)

type ItemService interface {
	FetchItems(string, string, string) (models.SearchResponse, error)
	SearchForItems(string) ([]models.SearchItem, error)
	PersistSearchResponse(models.SearchResponse) error
}

type OsrsService struct {
	itemRepo   persist.ItemRepository
	itemClient client.ItemClient
}

func NewOsrsService(repo persist.ItemRepository, client client.ItemClient) ItemService {
	return &OsrsService{repo, client}
}

func (s *OsrsService) FetchItems(category string, alpha string, page string) (models.SearchResponse, error) {
	resp, err := s.itemClient.FetchOsrsData(category, alpha, page)
	if err != nil {
		return models.SearchResponse{}, nil
	}
	return resp, nil
}

func (s *OsrsService) SearchForItems(itemName string) ([]models.SearchItem, error) {
	items, err := s.itemRepo.GetItem(itemName)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *OsrsService) PersistSearchResponse(respose models.SearchResponse) error {
	err := s.itemRepo.SaveItems(respose.Items)
	if err != nil {
		return err
	}
	return nil
}
