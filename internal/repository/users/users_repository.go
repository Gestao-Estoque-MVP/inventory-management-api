package users_repository

import (
	"context"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type IUsersRepository interface {
	CreateUser(user dto.UserCreateDTO) (*pgtype.UUID, error)
	GetUserByEmail(email string) (dto.GetUserByEmailDTO, error)
}

type UsersRepository struct {
	db *database.Queries
}

func NewRepositoryUsers(db *database.Queries) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) CreateUser(user dto.UserCreateDTO) (*pgtype.UUID, error) {
	id, err := r.db.CreateUser(context.Background(), database.CreateUserParams{
		Name:           pgtype.Text{String: user.Name, Valid: true},
		Email:          user.Email,
		Password:       pgtype.Text{String: user.Password, Valid: true},
		Document:       user.Document,
		MobilePhone:    pgtype.Text{String: user.MobilePhone, Valid: true},
		Active:         pgtype.Bool{Bool: true, Valid: true},
		RegisterToken:  pgtype.Text{String: "register_token", Valid: true},
		TokenExpiresAt: pgtype.Timestamp{Time: time.Now().Add(time.Hour * 24), Valid: true},
	})

	if err != nil {
		return &pgtype.UUID{}, err
	}

	return &id, nil
}

func (r *UsersRepository) GetUserByEmail(email string) (dto.GetUserByEmailDTO, error) {
	get, err := r.db.GetUserByEmail(context.Background(), email)
	if err != nil {
		return dto.GetUserByEmailDTO{}, err
	}

	return dto.GetUserByEmailDTO{
		ID:       get.ID,
		Email:    get.Email,
		Password: get.Password.String,
		Role:     get.RoleName,
	}, nil
}
