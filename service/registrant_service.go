package service

import (
	"context"
	"fmt"

	"github.com/pankop/event-porto/constants"
	"github.com/pankop/event-porto/dto"
	"github.com/pankop/event-porto/entity"
	"github.com/pankop/event-porto/repository"
)

type (
	RegistrantService interface {
		RegisterRegistrant(ctx context.Context, dto dto.RegistrantCreateRequest, reqP dto.PaymentRequest) (dto.RegistrantCreateResponse, error)
		// GetAllRegistrantWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.RegistrantPaginationResponse, error)
		// GetRegistrantById(ctx context.Context, id uuid.UUID) (dto.RegistrantCreateResponse, error)
		// GetRegistrantByEmail(ctx context.Context, email string) (dto.RegistrantCreateResponse, error)
		// UpdateRegistrant(ctx context.Context, id uuid.UUID, dto dto.RegistrantCreateRequest) (dto.RegistrantUpdateResponse, error)
	}
	registrantService struct {
		registrantRepo repository.RegistrantRepository
		paymentRepo    repository.PaymentRepository
	}
)

func NewRegistrantService(registrantRepo repository.RegistrantRepository, paymentRepo repository.PaymentRepository) RegistrantService {
	return &registrantService{
		registrantRepo: registrantRepo,
		paymentRepo:    paymentRepo,
	}
}

func (s *registrantService) RegisterRegistrant(ctx context.Context, req dto.RegistrantCreateRequest, reqP dto.PaymentRequest) (dto.RegistrantCreateResponse, error) {
	registrantDetail := req.RegistrantDetails[0]

	registrant := entity.EventRegistrants{
		Team_Name: registrantDetail.IdentitasTim.TeamName,
		School:    registrantDetail.IdentitasTim.School,
	}

	leaderDetail := entity.RegistrationDetails{
		Name:               registrantDetail.IdentitasKetua.Name,
		Email:              registrantDetail.IdentitasKetua.Email,
		PhoneNumber:        registrantDetail.IdentitasKetua.PhoneNumber,
		ImgIdentity:        registrantDetail.IdentitasKetua.ImgIdentity,
		ImgFollowInstagram: registrantDetail.IdentitasKetua.ImgFollowInstagram,
		Link_Twibbon:       registrantDetail.IdentitasKetua.LinkTwibbon,
		Role:               constants.Leader,
		Registrant_ID:      registrantDetail.IdentitasKetua.Registrant_ID,
	}

	// Memetakan detail anggota tim
	memberDetail := entity.RegistrationDetails{
		Name:               registrantDetail.IdentitasAnggota.Name,
		Email:              registrantDetail.IdentitasAnggota.Email,
		PhoneNumber:        registrantDetail.IdentitasAnggota.PhoneNumber,
		ImgIdentity:        registrantDetail.IdentitasAnggota.ImgIdentity,
		ImgFollowInstagram: registrantDetail.IdentitasAnggota.ImgFollowInstagram,
		Link_Twibbon:       registrantDetail.IdentitasAnggota.LinkTwibbon,
		Role:               constants.Member,
		Registrant_ID:      registrantDetail.IdentitasAnggota.Registrant_ID,
	}

	// Masukkan ketua dan anggota ke dalam field `Members` pada entitas registrant
	registrant.RegistrantDetails = []entity.RegistrationDetails{leaderDetail, memberDetail}

	payment := entity.Payments{
		Bank_ID:             reqP.Bank_ID,
		Bank_Transfer_From: reqP.Bank_Transfer_From,
		Name_Transfer_From: reqP.Name_Transfer_From,
		Amount:             reqP.Amount,
		Payment_Method:     reqP.Payment_Method,
		Payment_Proof:      reqP.Payment_Proof,
	}
	// Simpan registrant beserta detail anggota ke database
	registrantReg, err := s.registrantRepo.RegisterRegistrant(ctx, nil, registrant)
	if err != nil {
		return dto.RegistrantCreateResponse{}, dto.ErrCreateregistrant
	}

	paymentReg, err := s.paymentRepo.CreatePayment(ctx, nil, payment)
	if err != nil {
		return dto.RegistrantCreateResponse{}, dto.ErrCreateregistrant
	}

	// Contoh response DTO yang dihasilkan dari Service
	return dto.RegistrantCreateResponse{
		RegistrantDetails: []dto.RegistrantDetailResponse{
			{
				IdentitasTim: dto.IdentitasTimResponse{
					AccountID: fmt.Sprintf("%v", registrantReg.Registrant_ID), // ID Registrant
					Event:     "COMPETITION",                                  // Nama Event
					TeamName:  registrantReg.Team_Name,                        // Nama Tim
					School:    registrantReg.School,                           // Sekolah
					Status:    string(registrantReg.Status),                   // Status Registrasi
					Phase:     string(registrantReg.Comp_Status),              // Tahap Kompetisi
				},
				IdentitasKetua: dto.IdentitasPersonResponse{
					Name:               leaderDetail.Name,
					Email:              leaderDetail.Email,
					PhoneNumber:        leaderDetail.PhoneNumber,
					ImgIdentity:        leaderDetail.ImgIdentity,
					ImgFollowInstagram: leaderDetail.ImgFollowInstagram,
					LinkTwibbon:        leaderDetail.Link_Twibbon,
				},
				IdentitasAnggota: dto.IdentitasPersonResponse{
					Name:               memberDetail.Name,
					Email:              memberDetail.Email,
					PhoneNumber:        memberDetail.PhoneNumber,
					ImgIdentity:        memberDetail.ImgIdentity,
					ImgFollowInstagram: memberDetail.ImgFollowInstagram,
					LinkTwibbon:        memberDetail.Link_Twibbon,
				},
			},
		},
		PaymentDetails: dto.PaymentResponse{
			Payment_ID:         paymentReg.Payment_ID.String(),
			Registrant_ID:      paymentReg.Registrant_ID.String(),
			Bank_Transfer_From: paymentReg.Bank_Transfer_From,
			Name_Transfer_From: paymentReg.Name_Transfer_From,
			Amount:             float64(paymentReg.Amount),
			Payment_Method:     string(paymentReg.Payment_Method),
			Payment_Proof:      paymentReg.Payment_Proof,
			Status:             string(paymentReg.Status),
		},
	}, nil

}

// func (s *registrantService) GetAllRegistrantWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.RegistrantPaginationResponse, error) {

// 	dataWithPaginate, err := s.registrantRepo.GetAllRegistrantWithPagination(ctx, nil, req)
// 	if err != nil {
// 		return dto.RegistrantPaginationResponse{}, err
// 	}

// 	var datas []dto.RegistrantDetailResponse
// 	for _, registrant := range dataWithPaginate.Registrants {
// 		data := dto.RegistrantDetailResponse{
// 			Team_Name: registrant.IdentitasTim.Team_Name,
// 			School:    registrant.IdentitasTim.School,

// 		}

// 	}
// }
