package service

import (
	"context"
	"database/sql"

	"github.com/sungyo4869/portfolio/model"
)

type TagService struct {
	db *sql.DB
}

func NewTagService(db *sql.DB) *TagService {
	return &TagService{
		db: db,
	}
}

func (s *TagService) CreateTag(ctx context.Context) (*model.Tag, error) {
	var tag model.Tag

	return &tag, nil
}
func (s *TagService) ReadTag(ctx context.Context) ([]*model.Tag, error) {
	var tags []*model.Tag

	return tags, nil
}
func (s *TagService) UpdateTag(ctx context.Context) (*model.Tag, error) {
	var tag model.Tag

	return &tag, nil
}
func (s *TagService) DeleteTag(ctx context.Context) error {
	return nil
}

