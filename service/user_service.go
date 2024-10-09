package service

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pankop/event-porto/constants"
	"github.com/pankop/event-porto/dto"
	"github.com/pankop/event-porto/entity"
	"github.com/pankop/event-porto/helpers"
	"github.com/pankop/event-porto/repository"
	"github.com/pankop/event-porto/utils"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		//RegisterUserDetails(ctx context.Context, req dto.UserCreateDetailsRequest) (dto.UserDetailResponse, error)
		//GetUserDetails(ctx context.Context, accountID string) (dto.UserDetailResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
		GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error)
		SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)
		UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error)
		DeleteUser(ctx context.Context, userId string) error
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		RegisterAdmin(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

var (
	mu sync.Mutex
)

const (
	LOCAL_URL          = "http://localhost:3000"
	VERIFY_EMAIL_ROUTE = "register/verify_email"
)

func (s *userService) RegisterUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.Account{
		Email:      req.Email,
		Password:   req.Password,
		Role:       string(constants.User),
		IsVerified: false,
		AccountDetails: entity.AccountDetails{
			Name:         req.Name,
			Phone_Number: req.Phone_Number,
			Jenjang:      req.Jenjang,
		},
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	//buat nyimpen ke db untuk accountDetails
	

	draftEmail, err := makeVerificationEmail(userReg.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	err = utils.SendMail(userReg.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:              userReg.ID.String(),
		Email:           userReg.Email,
		Name:            userReg.AccountDetails.Name,
		Phone_Number:    userReg.AccountDetails.Phone_Number,
		Jenjang:         userReg.AccountDetails.Jenjang,
		IsEmailVerified: userReg.IsVerified,
		Role:            userReg.Role,
	}, nil
}

// func (s *userService) RegisterUserDetails(ctx context.Context, req dto.UserCreateDetailsRequest) (dto.UserDetailResponse, error) {
// 	userDetails := entity.AccountDetails{
// 		AccountID:    req.Account_ID,
// 		Name:         req.Name,
// 		Phone_Number: req.Phone_Number,
// 		Jenjang:      constants.Pelajar,
// 	}

// 	userReg, err := s.userRepo.RegisterUserDetails(ctx, nil, userDetails)
// 	if err != nil {
// 		return dto.UserDetailResponse{}, dto.ErrCreateUser
// 	}

// 	return dto.UserDetailResponse{
// 		ID:           userReg.ID.String(), // Ambil ID dari `AccountDetails`
// 		Name:         userReg.Name,
// 		Phone_Number: userReg.Phone_Number,
// 		Jenjang:      userReg.Jenjang,
// 		Account_ID:   userReg.AccountID,
// 	}, nil
// }

func makeVerificationEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := receiverEmail + "_" + expired
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	verifyLink := LOCAL_URL + "/" + VERIFY_EMAIL_ROUTE + "?token=" + token

	readHtml, err := os.ReadFile("utils/email-template/base_mail.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "pankop - Verification Email",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (s *userService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.ErrEmailNotFound
	}

	draftEmail, err := makeVerificationEmail(user.Email)
	if err != nil {
		return err
	}

	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error) {
	decryptedToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	if !strings.Contains(decryptedToken, "_") {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	decryptedTokenSplit := strings.Split(decryptedToken, "_")
	email := decryptedTokenSplit[0]
	expired := decryptedTokenSplit[1]

	now := time.Now()
	expiredTime, err := time.Parse("2006-01-02 15:04:05", expired)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	if expiredTime.Sub(now) < 0 {
		return dto.VerifyEmailResponse{
			Email:      email,
			IsVerified: false,
		}, dto.ErrTokenExpired
	}

	user, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
	}

	if user.IsVerified {
		return dto.VerifyEmailResponse{}, dto.ErrAccountAlreadyVerified
	}

	updatedUser, err := s.userRepo.UpdateUser(ctx, nil, entity.Account{
		ID:         user.ID,
		IsVerified: true,
	})
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
	}

	return dto.VerifyEmailResponse{
		Email:      email,
		IsVerified: updatedUser.IsVerified,
	}, nil
}

// user_service.go
// func (s *userService) GetUserDetails(ctx context.Context, accountID string) (dto.UserDetailResponse, error) {
// 	// Memanggil repository untuk mengambil `AccountDetails` dan `Account`
// 	userDetails, err := s.userRepo.GetUserDetails(ctx, accountID)
// 	if err != nil {
// 		return dto.UserDetailResponse{}, err
// 	}

// 	// Mengembalikan response yang mencakup informasi dari `AccountDetails` dan `Account`
// 	return dto.UserDetailResponse{
// 		ID:           userDetails.ID.String(), // Menggunakan Account_ID sebagai ID
// 		Name:         userDetails.Name,
// 		Phone_Number: userDetails.Phone_Number,
// 		Jenjang:      userDetails.Jenjang,
// 	}, nil
// }

func (s *userService) GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := s.userRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserPaginationResponse{}, err
	}

	var datas []dto.UserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.UserResponse{
			ID:              user.ID.String(),
			Name:            user.AccountDetails.Name,
			Email:           user.Email,
			Role:            user.Role,
			Phone_Number:    user.AccountDetails.Phone_Number,
			Jenjang:         user.AccountDetails.Jenjang,
			IsEmailVerified: user.IsVerified,
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:              user.ID.String(),
		Name:            user.AccountDetails.Name,
		Phone_Number:    user.AccountDetails.Phone_Number,
		Jenjang:         user.AccountDetails.Jenjang,
		Role:            user.Role,
		Email:           user.Email,
		IsEmailVerified: user.IsVerified,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	emails, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByEmail
	}

	//var userDetails entity.AccountDetails
	// userDetails, err := s.userRepo.GetUserDetails(ctx, emails.Account_ID)
	// if err != nil {
	// 	return dto.UserResponse{}, err
	// }
	return dto.UserResponse{
		ID:    emails.ID.String(),
		Email: emails.Email,
		//Name:            userDetails.Name,
		//Phone_Number:    userDetails.Phone_Number,
		Role:            emails.Role,
		IsEmailVerified: emails.IsVerified,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUserNotFound
	}

	data := entity.Account{
		ID: user.ID,
		AccountDetails: entity.AccountDetails{
			Name:         req.Name,
			Phone_Number: req.Phone_Number,
		},
		Email: req.Email,
	}

	userUpdate, err := s.userRepo.UpdateUser(ctx, nil, data)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUpdateUser
	}

	return dto.UserUpdateResponse{
		ID:              userUpdate.ID.String(),
		Name:            userUpdate.AccountDetails.Name,
		Phone_Number:    userUpdate.AccountDetails.Phone_Number,
		Role:            userUpdate.Role,
		Email:           userUpdate.Email,
		IsEmailVerified: user.IsVerified,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, userId string) error {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ErrUserNotFound
	}

	err = s.userRepo.DeleteUser(ctx, nil, user.ID.String())
	if err != nil {
		return dto.ErrDeleteUser
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	check, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	if !check.IsVerified {
		return dto.UserLoginResponse{}, dto.ErrAccountNotVerified
	}

	checkPassword, err := helpers.CheckPassword(check.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	token := s.jwtService.GenerateToken(check.ID.String(), "User")

	return dto.UserLoginResponse{
		Token: token,
		Role:  check.Role,
	}, nil
}

func (s *userService) RegisterAdmin(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.Account{
		Email:    req.Email,
		Password: req.Password,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID: userReg.ID.String(),
		//Name:            userReg.AccountDetails.Name,
		//Phone_Number:    userReg.AccountDetails.Phone_Number,
		Role:            string(constants.Admin),
		Email:           userReg.Email,
		IsEmailVerified: userReg.IsVerified,
	}, nil
}
