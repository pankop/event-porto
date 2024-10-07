package dto

import (
	"errors"

	"github.com/pankop/event-porto/constants"
	"github.com/pankop/event-porto/entity"
)

// Error messages
const (
	MESSAGE_FAILED_CREATE_REGISTRATION  = "failed registered for IoI Competition"
	MESSAGE_SUCCESS_CREATE_REGISTRATION = "successfully registered for IoI Competition"
	MESSAGE_FAILED_UPDATE_REGISTRATION  = "failed updated for IoI Competition"
	MESSAGE_SUCCESS_UPDATE_REGISTRATION = "successfully updated for IoI Competition"
)

// Error variables
var (
	ErrCreateregistrant             = errors.New("failed to create registrant")
	ErrGetAllregistrant             = errors.New("failed to get all registrant")
	ErrGetregistrantById            = errors.New("failed to get registrant by id")
	ErrGetregistrantByEmail         = errors.New("failed to get registrant by email")
	ErrUpdateregistrant             = errors.New("failed to update registrant")
	ErrregistrantNotFound           = errors.New("registrant not found")
	ErrRegistrantEmailAlreadyExists = errors.New("email already exist")
)

type (
	RegistrantCreateRequest struct {
		RegistrantDetails []RegistrantDetailRequest `json:"registrant_details" form:"registrant_details"`
		PaymentDetails    PaymentDetailRequest      `json:"payment_details" form:"payment_details"`
	}

	RegistrantUpdateRequest struct {
		RegistrantDetails []RegistrantDetailRequest `json:"registrant_details" form:"registrant_details"`
		PaymentDetails    PaymentDetailRequest      `json:"payment_details" form:"payment_details"`
	}

	// Response DTOs
	RegistrantCreateResponse struct {
		RegistrantDetails []RegistrantDetailResponse `json:"registrant_details"`
		PaymentDetails    PaymentDetailResponse      `json:"payment_details"`
	}

	RegistrantUpdateResponse struct {
		Status            constants.CompStatus       `json:"status"`
		RegistrantDetails []RegistrantDetailResponse `json:"registrant_details"`
		PaymentDetails    PaymentDetailResponse      `json:"payment_details"`
	}

	RegistrantPaginationResponse struct {
		Data []RegistrantCreateResponse `json:"data"`
		PaginationResponse
	}

	GetAllRegistrantRepositoryResponse struct {
		Registrant []entity.EventRegistrants
		PaginationResponse
	}

	// Detail Structure for Registrant
	RegistrantDetailRequest struct {
		IdentitasTim     IdentitasTimRequest    `json:"identitas_tim" form:"identitas_tim"`
		IdentitasKetua   IdentitasPersonRequest `json:"identitas_ketua" form:"identitas_ketua"`
		IdentitasAnggota IdentitasPersonRequest `json:"identitas_anggota" form:"identitas_anggota"`
	}

	RegistrantDetailResponse struct {
		IdentitasTim     IdentitasTimResponse    `json:"identitas_tim"`
		IdentitasKetua   IdentitasPersonResponse `json:"identitas_ketua"`
		IdentitasAnggota IdentitasPersonResponse `json:"identitas_anggota"`
	}

	IdentitasTimRequest struct {
		TeamName string `json:"team_name" form:"team_name" binding:"required"`
		School   string `json:"school" form:"school" binding:"required"`
	}

	IdentitasTimResponse struct {
		AccountID string `json:"account_id"`
		Event     string `json:"event"`
		TeamName  string `json:"team_name"`
		School    string `json:"school"`
		Status    string `json:"status"`
		Phase     string `json:"phase"`
	}

	IdentitasPersonRequest struct {
		Name               string `json:"name" form:"name" binding:"required"`
		Email              string `json:"email" form:"email" binding:"required"`
		PhoneNumber        string `json:"phone_number" form:"phone_number" binding:"required"`
		ImgIdentity        string `json:"img_identity" form:"img_identity" binding:"required"`
		ImgFollowInstagram string `json:"img_follow_instagram" form:"img_follow_instagram" binding:"required"`
		LinkTwibbon        string `json:"link_twibbon" form:"link_twibbon" binding:"required"`
		Role               string `json:"role" form:"role" binding:"required"`
	}

	IdentitasPersonResponse struct {
		Name               string `json:"name"`
		Email              string `json:"email"`
		PhoneNumber        string `json:"phone_number"`
		ImgIdentity        string `json:"img_identity"`
		ImgFollowInstagram string `json:"img_follow_instagram"`
		LinkTwibbon        string `json:"link_twibbon"`
	}

	// Payment Structures
	PaymentDetailRequest struct {
		BankID           *int64                  `json:"bank_id" form:"bank_id" example:"null"`
		BankTransferFrom string                  `json:"bank_transfer_from" form:"bank_transfer_from" binding:"required"`
		NameTransferFrom string                  `json:"name_transfer_from" form:"name_transfer_from" binding:"required"`
		FinalAmount      float64                 `json:"final_amount" form:"final_amount" binding:"required"`
		PaymentMethod    constants.PaymentMethod `json:"payment_method" form:"payment_method" binding:"required"`
		PaymentProof     string                  `json:"payment_proof" form:"payment_proof" binding:"required"`
	}

	PaymentDetailResponse struct {
		PaymentID        string  `json:"payment_id"`
		RegistrantID     string  `json:"registrant_id"`
		BankTransferFrom string  `json:"bank_transfer_from"`
		NameTransferFrom string  `json:"name_transfer_from"`
		FinalAmount      float64 `json:"final_amount"`
		PaymentMethod    string  `json:"payment_method"`
		PaymentProof     string  `json:"payment_proof"`
		Status           string  `json:"status"`
	}
)
