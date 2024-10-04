package repository

import (
	"context"

	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

type (
	ShortenLinkRepository interface {
		Create(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) (*entity.ShortenLink, error)
		Update(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) (*entity.ShortenLink, error)
		GetByOriginalLink(ctx context.Context, tx *gorm.DB, originalLink string) (*entity.ShortenLink, error)
		GetByShortenLink(ctx context.Context, tx *gorm.DB, shortenLink string) (*entity.ShortenLink, error)
		GetAll(ctx context.Context, tx *gorm.DB) ([]entity.ShortenLink, error)
		Delete(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) error
	}

	shortenLinkRepository struct {
		db *gorm.DB
	}
)

func NewShortenLinkRepository(db *gorm.DB) ShortenLinkRepository {
	return &shortenLinkRepository{
		db: db,
	}
}

func (r *shortenLinkRepository) Create(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) (*entity.ShortenLink, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&shortenLink).Error; err != nil {
		return &entity.ShortenLink{}, err
	}

	return shortenLink, nil
}

func (r *shortenLinkRepository) Update(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) (*entity.ShortenLink, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&shortenLink).Error; err != nil {
		return &entity.ShortenLink{}, err
	}

	return shortenLink, nil
}

func (r *shortenLinkRepository) GetByOriginalLink(ctx context.Context, tx *gorm.DB, originalLink string) (*entity.ShortenLink, error) {
	if tx == nil {
		tx = r.db
	}

	var shortenLink entity.ShortenLink
	if err := tx.WithContext(ctx).Where("original_link = ?", originalLink).Take(&shortenLink).Error; err != nil {
		return &entity.ShortenLink{}, err
	}

	return &shortenLink, nil
}

func (r *shortenLinkRepository) GetByShortenLink(ctx context.Context, tx *gorm.DB, shortenLink string) (*entity.ShortenLink, error) {
	if tx == nil {
		tx = r.db
	}

	var shorten_Link entity.ShortenLink
	if err := tx.WithContext(ctx).Where("shorten_link = ?", shortenLink).Take(&shorten_Link).Error; err != nil {
		return &entity.ShortenLink{}, err
	}

	return &shorten_Link, nil
}

func (r *shortenLinkRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]entity.ShortenLink, error) {
	if tx == nil {
		tx = r.db
	}

	var shortenLinks []entity.ShortenLink
	if err := tx.WithContext(ctx).Find(&shortenLinks).Error; err != nil {
		return []entity.ShortenLink{}, err
	}

	return shortenLinks, nil
}

func (r *shortenLinkRepository) Delete(ctx context.Context, tx *gorm.DB, shortenLink *entity.ShortenLink) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&shortenLink).Error; err != nil {
		return err
	}

	return nil
}
