package repository

import (
	"context"
	"math"

	"github.com/pankop/event-porto/dto"
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

type (
	PaymentRepository interface {
		CreatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payments) (entity.Payments, error)
		GetAllPaymentWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPaymentResponse, error)
		GetPaymentById(ctx context.Context, tx *gorm.DB, paymentId string) (*entity.Payments, error)
		UpdatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payments) (entity.Payments, error)
		DeletePayment(ctx context.Context, tx *gorm.DB, paymentId string) error
	}
	paymentRepository struct {
		db *gorm.DB
	}
)

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payments) (entity.Payments, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&payment).Error; err != nil {
		return entity.Payments{}, err
	}

	return payment, nil
}

func (r *paymentRepository) GetPaymentById(ctx context.Context, tx *gorm.DB, paymentId string) (*entity.Payments, error) {
	if tx == nil {
		tx = r.db
	}
	var payment entity.Payments
	if err := tx.WithContext(ctx).Where("id = ?", paymentId).Take(&payment).Error; err != nil {
		return &entity.Payments{}, err
	}

	return &payment, nil
}

func (r *paymentRepository) GetAllPaymentWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPaymentResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var payments []entity.Payments
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Payments{}).Count(&count).Error; err != nil {
		return dto.GetAllPaymentResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&payments).Error; err != nil {
		return dto.GetAllPaymentResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllPaymentResponse{
		Payments: payments,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}

func (r *paymentRepository) UpdatePayment(ctx context.Context, tx *gorm.DB, payment entity.Payments) (entity.Payments, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&payment).Error; err != nil {
		return entity.Payments{}, err
	}

	return payment, nil
}

func (r *paymentRepository) DeletePayment(ctx context.Context, tx *gorm.DB, paymentId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Payments{}, "id = ?", paymentId).Error; err != nil {
		return err
	}

	return nil
}
