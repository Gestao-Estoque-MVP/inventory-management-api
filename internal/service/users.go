package service

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
	"github.com/diogoX451/inventory-management-api/pkg/convert"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	userRepo  repository.IUserRepository
	role      repository.IRBCA
	email     *EmailService
	s3Service *S3Service
}

func NewServiceUser(userRepo repository.IUserRepository, role repository.IRBCA, email *EmailService, s3 *S3Service) *UserService {
	return &UserService{userRepo: userRepo, role: role, email: email, s3Service: s3}
}

func (s *UserService) CreateTenant(tenant *database.Tenant) (*database.Tenant, error) {
	id, _ := uuid.NewV4()
	params := &database.Tenant{
		ID:        pgtype.UUID{Bytes: id, Valid: true},
		Name:      tenant.Name,
		TaxCode:   tenant.TaxCode,
		Type:      tenant.Type,
		CreatedAt: pgtype.Timestamptz{Time: time.Now().Local(), Valid: true},
	}

	create, err := s.userRepo.CreateTenant(params)

	if err != nil {
		return nil, err
	}

	return create, err
}

func (us *UserService) CreatePreUser(user *database.CreatePreRegisterUserParams) (*pgtype.UUID, error) {
	token, _ := token.GeneratedToken()
	id, _ := uuid.NewV4()
	phoneId, _ := uuid.NewV4()
	role, _ := us.role.GetRole("users")

	params := &database.CreatePreRegisterUserParams{
		ID:             pgtype.UUID{Bytes: id, Valid: true},
		Name:           user.Name,
		Email:          user.Email,
		Status:         database.UserStatusPreUsers,
		TenantID:       user.TenantID,
		RegisterToken:  pgtype.Text{String: token, Valid: true},
		TokenExpiresAt: pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour), Valid: true},
		CreatedAt:      pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
		ID_2:           pgtype.UUID{Bytes: phoneId, Valid: true},
		Type:           user.Type,
		Number:         user.Number,
		IsPrimary:      user.IsPrimary,
		CreatedAt_2:    pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	}

	createUser, err := us.userRepo.CreatePreUser(params, role.ID.Bytes)

	if err != nil {
		log.Printf("Erro ao criar usuário: %v\n", err)
		return nil, err
	}

	// detail := &EmailDetails{
	// 	To:         []string{createUser.Email},
	// 	Subject:    "Pré-Cadastro no SwiftStock",
	// 	TemplateID: "947cd590-5b82-4e1c-a7db-c80f6534168b",
	// }

	// err = us.email.SendEmail(detail, "one")
	// if err != nil {
	// 	log.Printf("error sending %v "+createUser.Email+":", err)
	// }

	return createUser, nil
}

func (us *UserService) CreateImageUser(image *database.CreateImageUserParams) (*pgtype.UUID, error) {
	create, err := us.userRepo.CreateImageUser(image)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (us *UserService) CompleteRegisterUser(registerToken string, user *database.CompleteRegisterUserParams, image io.Reader) (*database.CompleteRegisterUserRow, error) {

	verifyUser, err := us.userRepo.GetUserRegisterToken(registerToken)

	if err != nil {
		log.Printf("Erro ao buscar usuário: %v\n", err)
		return nil, fmt.Errorf("no user found with register token %s", registerToken)
	}

	if image != nil {
		id, _ := uuid.NewV4()
		upload, err := us.s3Service.UploadImageS3(image, convert.UUIDToString(verifyUser.ID))
		if err != nil {
			return nil, err
		}
		us.CreateImageUser(&database.CreateImageUserParams{
			ID:          pgtype.UUID{Bytes: id, Valid: true},
			Description: verifyUser.Name,
			Url:         pgtype.Text{String: upload, Valid: true},
			CreatedAt:   pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).Local(), Valid: true},
			ID_2:        verifyUser.ID,
		})
	}

	user.Status = database.UserStatusActive
	updateUser, err := us.userRepo.CreateCompleteUser(verifyUser.RegisterToken.String, user)

	if err != nil {
		log.Printf("Erro ao criar usuário completo %v\n", err)
		return nil, err
	}

	return updateUser, nil
}

func (us *UserService) CreateCompanyUser(user *database.User, roleId [][16]byte) (*database.CreateCompanyUsersRow, error) {
	token, _ := token.GeneratedToken()
	user.RegisterToken = pgtype.Text{String: token, Valid: true}
	user.TokenExpiresAt = pgtype.Timestamp{Time: time.Now().Add(1 * time.Hour).Local(), Valid: true}
	user.Status = database.UserStatusActive

	create, err := us.userRepo.CreateCompanyUser(user, roleId)

	if err != nil {
		return nil, err
	}

	return create, nil

}

func (us *UserService) UpdateUser(id [16]byte, user *database.UpdateUserParams) error {
	err := us.userRepo.UpdateUser(id, user)

	if err != nil {
		log.Printf("Erro ao atualizar usuário completo %v\n", err)
		return err
	}

	return err
}

func (us *UserService) GetUsers() ([]*database.User, error) {
	list, err := us.userRepo.GetUsers()

	if err != nil {
		log.Printf("Erro ao trazer a listar usuário completo %v\n", err)
		return nil, err
	}

	return list, nil
}

func (us UserService) GetUser(id pgtype.UUID) (*database.GetUserRow, error) {
	get, err := us.userRepo.GetUser(id)

	if err != nil {
		log.Printf("Erro ao trazer usuário: %v\n", err)
		return nil, err
	}

	if get.Image.Url.String != "" {
		url, err := us.s3Service.GetUrlS3(get.Image.Url.String)
		if err != nil {
			log.Printf("Erro ao obter URL do S3: %v\n", err)
			return nil, err
		}
		get.Image.Url = pgtype.Text{String: url, Valid: true}
	}

	return get, nil
}

func (us *UserService) GetUserByEmail(email string) (*database.GetEmailRow, error) {
	get, err := us.userRepo.GetUserByEmail(email)

	if err != nil {
		log.Printf("Erro ao trazer usuário %v\n", err)
		return nil, err
	}

	return get, nil
}

func (us *UserService) VerifyToken(token string) bool {
	find, err := us.userRepo.VerifyToken(token)

	if err != nil {
		return false
	}

	return find.TokenExpiresAt.Time == time.Now().Local().Add(time.Hour*2)

}
