package service

import (
	"context"
	"database/sql"

	"github.com/sungyo4869/portfolio/model"
)

type PictureService struct {
	db *sql.DB
}

func NewPictureService(db *sql.DB) *PictureService {
	return &PictureService{
		db: db,
	}
}

func (s *PictureService) CreatePicture(ctx context.Context) ([]*model.Picture, error) {
	var pictures []*model.Picture

	return pictures, nil
}
func (s *PictureService) ReadPicture(ctx context.Context) ([]*model.Picture, error) {
	var pictures []*model.Picture

	return pictures, nil
}
func (s *PictureService) UpdatePicture(ctx context.Context) ([]*model.Picture, error) {
	var pictures []*model.Picture

	return pictures, nil
}
func (s *PictureService) DeletePicture(ctx context.Context) error {
	return nil
}

