package repository

import (
	"context"
	"math"

	"github.com/pankop/event-porto/dto"
	"github.com/pankop/event-porto/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.Account) (entity.Account, error)
		RegisterUserDetails(ctx context.Context, tx *gorm.DB, user entity.AccountDetails) (entity.AccountDetails, error)
		GetUserDetails(ctx context.Context, accountID string) (entity.AccountDetails, error)
		GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllUserRepositoryResponse, error)
		GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.Account, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Account, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Account, bool, error)
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.Account) (entity.Account, error)
		DeleteUser(ctx context.Context, tx *gorm.DB, userId string) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user entity.Account) (entity.Account, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.Account{}, err
	}

	return user, nil
}
	
func (r *userRepository) RegisterUserDetails(ctx context.Context, tx *gorm.DB, user entity.AccountDetails) (entity.AccountDetails, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.AccountDetails{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllUserRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var users []entity.Account
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Account{}).Count(&count).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&users).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllUserRepositoryResponse{
		Users: users,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *userRepository) GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.Account, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.Account
	if err := tx.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		return entity.Account{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Account, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.Account
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.Account{}, err
	}

	return user, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Account, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.Account
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.Account{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.Account) (entity.Account, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.Account{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, tx *gorm.DB, userId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Account{}, "id = ?", userId).Error; err != nil {
		return err
	}

	return nil
}

// user_repository.go
func (r *userRepository) GetUserDetails(ctx context.Context, accountID string) (entity.AccountDetails, error) {
	// Mengambil data dari entity `AccountDetails`
	var userDetails entity.AccountDetails
	err := r.db.Where("account_id = ?", accountID).First(&userDetails).Error
	if err != nil {
		return entity.AccountDetails{}, err
	}

	// Mengambil data dari entity `Account` berdasarkan `accountID`

	return userDetails, nil
}
