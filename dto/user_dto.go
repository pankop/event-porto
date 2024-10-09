package dto

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pankop/event-porto/constants"
	"github.com/pankop/event-porto/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL            = "failed verify email"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
	MESSAGE_SUCCESS_REGISTER_USER_DETAIL    = "success register user detail"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotAdmin           = errors.New("user not admin")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrAccountDetailsNotFound = errors.New("account details not found")
)

type (
	UserCreateRequest struct {
		Email        string            `json:"email" form:"email"`
		Password     string            `json:"password" form:"password"`
		Name         string            `json:"name" form:"name"`
		Phone_Number string            `json:"phone_number" form:"phone_number"`
		Jenjang      constants.Jenjang `json:"jenjang" form:"jenjang"`
	}

	UserResponse struct {
		ID              string            `json:"id"`
		Email           string            `json:"email"`
		IsEmailVerified bool              `json:"is_verified"`
		Role            string            `json:"role"`
		Name            string            `json:"name" form:"name"`
		Phone_Number    string            `json:"phone_number" form:"phone_number"`
		Jenjang         constants.Jenjang `json:"jenjang" form:"jenjang"`
	}

	// UserCreateDetailsRequest struct {
	// 	Account_ID   string `json:"account_id"`
	// 	Name         string `json:"name"`
	// 	Phone_Number string `json:"phone_number"`
	// 	Jenjang      string `json:"jenjang"`
	// }

	// dto/user_response.go
	// UserDetailResponse struct {
	// 	ID           string            `json:"id"`
	// 	Name         string            `json:"name"`
	// 	Phone_Number string            `json:"phone_number"`
	// 	Jenjang      constants.Jenjang `json:"jenjang"`
	// 	Account_ID   string            `json:"account_id"`
	// 	// Email           string            `json:"email"`             // Tambahkan field email
	// 	// IsEmailVerified bool              `json:"is_email_verified"` // Status verifikasi email
	// 	// Role            string            `json:"role"`
	// }

	UserPaginationResponse struct {
		Data []UserResponse `json:"data"`
		PaginationResponse
	}

	GetAllUserRepositoryResponse struct {
		Users []entity.Account
		PaginationResponse
	}

	UserUpdateRequest struct {
		Name         string `json:"name" form:"name"`
		Phone_Number string `json:"telp_number" form:"telp_number"`
		Email        string `json:"email" form:"email"`
	}

	UserUpdateResponse struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Phone_Number    string `json:"telp_number"`
		Role            string `json:"role"`
		Email           string `json:"email"`
		IsEmailVerified bool   `json:"is_verified"`
	}

	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		Role  string    `json:"role"`
		Token string    `json:"token"`
	}

	UpdateStatusIsVerifiedRequest struct {
		UserId     string `json:"user_id" form:"user_id" binding:"required"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}
)
