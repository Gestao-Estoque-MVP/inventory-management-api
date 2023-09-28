package repository

import (
	"context"
	"log"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreatePreUser(user *database.CreatePreRegisterUserParams, roleId [][16]byte) (*pgtype.UUID, error)
	CreateCompleteUser(token string, user *database.CompleteRegisterUserParams) (*database.CompleteRegisterUserRow, error)
	CreateCompanyUser(user *database.User, roleId [][16]byte) (*database.CreateCompanyUsersRow, error)
	CreateUserPhones(*database.UserPhone) (*database.CreateUserPhonesRow, error)
	CreateTenant(tenant *database.Tenant) (*database.Tenant, error)
	CreateImageUser(image *database.CreateImageUserParams) (*pgtype.UUID, error)
	UpdateUser(id [16]byte, user *database.UpdateUserParams) error
	DeleteUser(id [16]byte) (bool, error)
	GetUser(id pgtype.UUID) (*database.GetUserRow, error)
	GetUsers() ([]*database.User, error)
	GetUserByEmail(email string) (*database.GetEmailRow, error)
	GetUserRegisterToken(token string) (*database.GetUserRegisterTokenRow, error)
	VerifyToken(token string) (*database.GetTokenPreRegisterRow, error)
	GetUsersByEmail() ([]*string, error)
	GetContacts() ([]*string, error)
	GetContact(email string) (*database.GetUserContactEmailRow, error)
}

type UserRepository struct {
	DB *database.Queries
}

func NewRepositoryUser(db *database.Queries) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateTenant(tenant *database.Tenant) (*database.Tenant, error) {
	create, err := r.DB.CreateTenant(context.Background(), database.CreateTenantParams{
		ID:        tenant.ID,
		Name:      tenant.Name,
		TaxCode:   tenant.TaxCode,
		Type:      tenant.Type,
		CreatedAt: tenant.CreatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (i *UserRepository) CreatePreUser(user *database.CreatePreRegisterUserParams, roleId [][16]byte) (*pgtype.UUID, error) {
	log.Printf("Creating user %", user.TenantID)
	create, err := i.DB.CreatePreRegisterUser(context.Background(), database.CreatePreRegisterUserParams{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		TenantID:       user.TenantID,
		RegisterToken:  user.RegisterToken,
		TokenExpiresAt: user.TokenExpiresAt,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt,
		ID_2:           user.ID_2,
		Type:           user.Type,
		Number:         user.Number,
		IsPrimary:      user.IsPrimary,
		CreatedAt_2:    user.CreatedAt_2,
	})

	if err != nil {
		return nil, err
	}

	for j := range roleId {
		_, err := i.DB.CreateUsersRoles(context.Background(), database.CreateUsersRolesParams{
			UserID: pgtype.UUID{Bytes: create.Bytes, Valid: true},
			RoleID: pgtype.UUID{Bytes: roleId[j], Valid: true},
		})
		if err != nil {
			log.Printf("Error created user role: %v", err)
			continue
		}
	}

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (i *UserRepository) CreateUserPhones(user *database.UserPhone) (*database.CreateUserPhonesRow, error) {
	create, err := i.DB.CreateUserPhones(context.Background(), database.CreateUserPhonesParams{
		ID:        user.ID,
		UserID:    user.UserID,
		Type:      user.Type,
		Number:    user.Number,
		IsPrimary: user.IsPrimary,
		CreatedAt: user.CreatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (i *UserRepository) CreateCompleteUser(token string, user *database.CompleteRegisterUserParams) (*database.CompleteRegisterUserRow, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	updateUser, err := i.DB.CompleteRegisterUser(context.Background(), database.CompleteRegisterUserParams{
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
		Password:       pgtype.Text{String: string(bytes), Valid: true},
		RegisterToken:  pgtype.Text{String: token, Valid: true},
		Status:         user.Status,
		UpdatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &database.CompleteRegisterUserRow{
		Name: updateUser.Name,
	}, nil
}

func (i *UserRepository) CreateImageUser(image *database.CreateImageUserParams) (*pgtype.UUID, error) {
	create, err := i.DB.CreateImageUser(context.Background(), database.CreateImageUserParams{
		ID:          image.ID,
		Url:         image.Url,
		Description: image.Description,
		CreatedAt:   image.CreatedAt,
		ID_2:        image.ID_2,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (i *UserRepository) CreateCompanyUser(user *database.User, roleId [][16]byte) (*database.CreateCompanyUsersRow, error) {
	create, err := i.DB.CreateCompanyUsers(context.Background(), database.CreateCompanyUsersParams{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Status:         user.Status,
		RegisterToken:  user.RegisterToken,
		TokenExpiresAt: user.TokenExpiresAt,
		CreatedAt:      user.CreatedAt,
		TenantID:       user.TenantID,
	})

	go func() {
		for j := range roleId {
			_, err := i.DB.CreateUsersRoles(context.Background(), database.CreateUsersRolesParams{
				UserID: create.ID,
				RoleID: pgtype.UUID{Bytes: roleId[j], Valid: true},
			})
			if err != nil {
				log.Printf("Error created user role: %v", err)
				continue
			}
		}
	}()

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (i *UserRepository) UpdateUser(id [16]byte, user *database.UpdateUserParams) error {

	err := i.DB.UpdateUser(context.Background(), database.UpdateUserParams{
		ID:             pgtype.UUID{Bytes: id, Valid: true},
		Name:           user.Name,
		Email:          user.Email,
		DocumentType:   user.DocumentType,
		DocumentNumber: user.DocumentNumber,
	})

	if err != nil {
		return err
	}

	return err
}

func (i *UserRepository) DeleteUser(id [16]byte) (bool, error) {
	_, err := i.DB.DeleteUser(context.Background(), pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *UserRepository) GetUser(id pgtype.UUID) (*database.GetUserRow, error) {
	get, err := i.DB.GetUser(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) GetUsers() ([]*database.User, error) {
	list, err := i.DB.ListUsers(context.Background())

	if err != nil {
		log.Printf("Error getting users %v", err)
		return nil, err
	}

	ptrList := make([]*database.User, len(list))
	for i := range list {
		user := list[i]
		ptrList[i] = &user
	}

	return ptrList, nil
}

func (i *UserRepository) GetUserByEmail(email string) (*database.GetEmailRow, error) {
	get, err := i.DB.GetEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) GetUserRegisterToken(token string) (*database.GetUserRegisterTokenRow, error) {
	get, err := i.DB.GetUserRegisterToken(context.Background(), pgtype.Text{String: token, Valid: true})

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (i *UserRepository) VerifyToken(token string) (*database.GetTokenPreRegisterRow, error) {
	find, err := i.DB.GetTokenPreRegister(context.Background(), pgtype.Text{
		String: token,
		Valid:  true,
	})

	if err != nil {
		return nil, err
	}

	return &find, nil
}

func (i *UserRepository) GetUsersByEmail() ([]*string, error) {
	get, err := i.DB.GetUsersWithEmail(context.Background())

	if err != nil {
		return nil, err
	}

	list := make([]*string, len(get))

	for i := range get {
		user := get[i]
		list[i] = &user
	}

	return list, nil
}

func (i *UserRepository) GetContacts() ([]*string, error) {
	list, err := i.DB.GetUsersContact(context.Background())
	if err != nil {
		return nil, err
	}

	contact := make([]*string, len(list))
	for i := range list {
		user := list[i]
		contact[i] = &user
	}

	return contact, nil
}

func (i *UserRepository) GetContact(email string) (*database.GetUserContactEmailRow, error) {
	list, err := i.DB.GetUserContactEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
