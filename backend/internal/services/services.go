package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"context"
)

type Service struct {
	database database.Database
}

func NewService(db database.Database) *Service {
	return &Service{database: db}
}

func (s *Service) GetStatuses(ctx context.Context) ([]models.PingStatus, error) {
	var statuses []models.PingStatus
	statuses, err := s.database.SelectStatusesQuery(ctx)
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (s *Service) UpdateStatus(ctx context.Context, ping models.PingStatus) error {
	err := s.database.UpdateStatusQuery(ctx, ping)
	return err
}
