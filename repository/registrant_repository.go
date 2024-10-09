package repository

import (
	"context"
	"math"

	"github.com/pankop/event-porto/dto"
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

type (
	RegistrantRepository interface {
		RegisterRegistrant(ctx context.Context, tx *gorm.DB, registrant entity.EventRegistrants) (entity.EventRegistrants, error)
		GetAllRegistrantWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllRegistrantRepositoryResponse, error)
		GetRegistrantById(ctx context.Context, tx *gorm.DB, registrantId string) (entity.EventRegistrants, error)
		GetRegistrantByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.EventRegistrants, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.EventRegistrants, bool, error)
		UpdateRegistrant(ctx context.Context, tx *gorm.DB, registrant entity.EventRegistrants) (entity.EventRegistrants, error)
		DeleteRegistrant(ctx context.Context, tx *gorm.DB, registrantId string) error
	}

	registrantRepository struct {
		db *gorm.DB
	}
)

func NewRegistrantRepository(db *gorm.DB) RegistrantRepository {
	return &registrantRepository{
		db: db,
	}
}

func (r *registrantRepository) RegisterRegistrant(ctx context.Context, tx *gorm.DB, registrant entity.EventRegistrants) (entity.EventRegistrants, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&registrant).Error; err != nil {
		return entity.EventRegistrants{}, err
	}

	return registrant, nil
}

func (r *registrantRepository) GetAllRegistrantWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllRegistrantRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var registrants []entity.EventRegistrants
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.EventRegistrants{}).Count(&count).Error; err != nil {
		return dto.GetAllRegistrantRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&registrants).Error; err != nil {
		return dto.GetAllRegistrantRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllRegistrantRepositoryResponse{
		Registrants: registrants,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}

func (r *registrantRepository) GetRegistrantById(ctx context.Context, tx *gorm.DB, registrantId string) (entity.EventRegistrants, error) {
	if tx == nil {
		tx = r.db
	}

	var registrant entity.EventRegistrants
	if err := tx.WithContext(ctx).Where("id = ?", registrantId).Take(&registrant).Error; err != nil {
		return entity.EventRegistrants{}, err
	}

	return registrant, nil
}

func (r *registrantRepository) GetRegistrantByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.EventRegistrants, error) {
	if tx == nil {
		tx = r.db
	}

	var registrant entity.EventRegistrants
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&registrant).Error; err != nil {
		return entity.EventRegistrants{}, err
	}

	return registrant, nil
}

func (r *registrantRepository) UpdateRegistrant(ctx context.Context, tx *gorm.DB, registrant entity.EventRegistrants) (entity.EventRegistrants, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&registrant).Error; err != nil {
		return entity.EventRegistrants{}, err
	}

	return registrant, nil
}

func (r *registrantRepository) DeleteRegistrant(ctx context.Context, tx *gorm.DB, registrantId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.EventRegistrants{}, "id = ?", registrantId).Error; err != nil {
		return err
	}

	return nil
}

func (r *registrantRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.EventRegistrants, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var registrant entity.EventRegistrants
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&registrant).Error; err != nil {
		return entity.EventRegistrants{}, false, err
	}

	return registrant, true, nil
}
