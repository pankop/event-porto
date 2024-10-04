package dto

import (
	"errors"

	"github.com/google/uuid"
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
	ErrCreateregistrant     = errors.New("failed to create registrant")
	ErrGetAllregistrant     = errors.New("failed to get all registrant")
	ErrGetregistrantById    = errors.New("failed to get registrant by id")
	ErrGetregistrantByEmail = errors.New("failed to get registrant by email")
	ErrUpdateregistrant     = errors.New("failed to update registrant")
	ErrregistrantNotFound   = errors.New("registrant not found")
)

type (
	RegistrantCreateRequest struct {
		CompetitionType   constants.CompetitionType `json:"competition_type" form:"competition_type" binding:"required"`
		RegistrantDetails []RegistrantDetail        `json:"registrant_details" form:"registrant_details"`
		PaymentDetails    PaymentDetail             `json:"payment_details" form:"payment_details"`
	}
	RegistrantCreateResopnse struct {
		ID                uuid.UUID            `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		RegistrantId      string               `json:"registrant_id"`
		Status            constants.CompStatus `json:"status" `
		CompetitionType   constants.CompetitionType
		RegistrantDetails []RegistrantDetail
		PaymentDetails    PaymentDetail
	}
	RegistrantPaginationResponse struct {
		Data []RegistrantCreateResopnse `json:"data"`
		PaginationResponse
	}

	GetAllRegistrantRepositoryResponse struct {
		Registrant []entity.EventRegistrants
		PaginationResponse
	}

	RegistrantDetail struct {
		IdentitasTim     IdentitasTim     `json:"identitas_tim" form:"identitas_tim"`
		IdentitasKetua   IdentitasKetua   `json:"identitas_ketua" form:"identitas_ketua"`
		IdentitasAnggota IdentitasAnggota `json:"identitas_anggota" form:"identitas_anggota"`
	}

	IdentitasTim struct {
		Team_Name string `json:"team_name" form:"team_name" binding:"required" `
		School    string `json:"school" form:"school" binding:"required" `
	}

	IdentitasKetua struct {
		Name               string `json:"name" form:"name" binding:"required" `
		Email              string `json:"email" form:"email" binding:"required" `
		Phone_Number       string `json:"phone_number" form:"phone_number" binding:"required"`
		ImgIdentity        string `json:"img_identity" form:"img_identity" binding:"required" `
		ImgFollowInstagram string `json:"img_follow_instagram" form:"img_follow_instagram" binding:"required" `
		Link_Twibbon       string `json:"link_twibbon" form:"link_twibbon" binding:"required" `
	}

	IdentitasAnggota struct {
		Name               string `json:"name" form:"name" binding:"required" `
		Email              string `json:"email" form:"email" binding:"required" `
		Phone_Number       string `json:"phone_number" form:"phone_number" binding:"required" `
		ImgIdentity        string `json:"img_identity" form:"img_identity" binding:"required" `
		ImgFollowInstagram string `json:"img_follow_instagram" form:"img_follow_instagram" binding:"required"`
		Link_Twibbon       string `json:"link_twibbon" form:"link_twibbon" binding:"required"`
	}

	PaymentDetail struct {
		BankID           *int64                  `json:"bank_id" form:"bank_id" example:"null"`
		BankTransferFrom string                  `json:"bank" form:"bank_transfer_from" binding:"required"`
		NameTransferFrom string                  `json:"name" form:"name_transfer_from" binding:"required"`
		FinalAmount      float64                 `json:"final_amount" form:"final_amount" binding:"required"`
		PaymentMethod    constants.PaymentMethod `json:"payment_method" form:"payment_method" binding:"required" `
		PaymentProof     string                  `json:"payment_proof" form:"payment_proof" binding:"required"`
	}

	RegistrantUpdateRequest struct {
		CompetitionType   constants.CompetitionType `json:"competition_type" form:"competition_type" binding:"required"`
		RegistrantDetails []RegistrantDetail        `json:"registrant_details" form:"registrant_details"`
		PaymentDetails    PaymentDetail             `json:"payment_details" form:"payment_details"`
	}

	RegistrantUpdateResponse struct {
		ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		RegistrantId      string    `json:"registrant_id"`
		Status            string    `json:"status" example:"pending"`
		CompetitionType   constants.CompetitionType
		registrantDetails []RegistrantDetail
		PaymentDetails    PaymentDetail
	}
)
