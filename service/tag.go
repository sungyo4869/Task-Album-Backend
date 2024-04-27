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

func (s *TagService) CreateTag(ctx context.Context, cardID int64, label string) (*model.Tag, error) {
	const (
		insert  = `INSERT INTO tags(card_id, label) VALUES(?, ?)`
		confirm = `SELECT id, card_id, label FROM tags WHERE id = ?`
	)

	result, err := s.db.ExecContext(ctx, insert, cardID, label)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var tag model.Tag
	row := s.db.QueryRowContext(ctx, confirm, id)
	err = row.Scan(
		&tag.ID, 
		&tag.CardID, 
		&tag.Label,
	)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
func (s *TagService) ReadTag(ctx context.Context) ([]*model.Tag, error) {
	const query = `SELECT id, card_id, label FROM tags`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*model.Tag
	for rows.Next() {
		var tag model.Tag
		err := rows.Scan(
			&tag.ID, 
			&tag.CardID,
			&tag.Label,
		)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagService) DeleteTag(ctx context.Context, cardID, tagID int64) error {
	const query = `DELETE FROM tags WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, tagID)
	if err != nil {
		return err
	}
	return nil
}

