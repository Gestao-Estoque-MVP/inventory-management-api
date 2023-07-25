package service_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository/mock"
	"github.com/diogoX451/inventory-management-api/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePreUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	// Definir expectativas para as chamadas de mock
	user := &database.User{
		Name:     "Test",
		Email:    "test@email.com",
		Status:   "Active",
		TenantID: "123",
	}

	role := &database.Role{
		ID:   "users",
		Name: "Users",
	}

	mockRoleRepo.EXPECT().GetRole("users").Return(role, nil)
	mockUserRepo.EXPECT().CreatePreUser(gomock.Any()).Return(user, nil)

	// Criar o UserService
	us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

	// Testar a função CreatePreUser
	result, err := us.CreatePreUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user, result)
}

func TestCompleteRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	registerToken := "token123"
	user := &database.User{
		Name:     "Test",
		Email:    "test@email.com",
		Status:   "Active",
		TenantID: "123",
	}

	verifyUser := &database.User{
		RegisterToken: sql.NullString{String: registerToken, Valid: true},
	}

	completedUser := &database.CompleteRegisterUserRow{}

	// Definir expectativas para as chamadas de mock
	mockUserRepo.EXPECT().GetUserRegisterToken(registerToken).Return(verifyUser, nil)
	mockUserRepo.EXPECT().CreateCompleteUser(gomock.Eq(registerToken), gomock.Any()).Return(completedUser, nil)

	// Criar o UserService
	us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

	// Testar a função CompleteRegisterUser
	result, err := us.CompleteRegisterUser(registerToken, user)
	assert.Nil(t, err)
	assert.Equal(t, completedUser, result)
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	// Crie uma lista de usuários que você espera que seja retornada pelo seu repositório
	users := []*database.User{
		{
			Name:     "Test",
			Email:    "test@email.com",
			Status:   "Active",
			TenantID: "123",
		},
		{
			Name:     "Test 2",
			Email:    "test2@email.com",
			Status:   "Inactive",
			TenantID: "123",
		},
	}

	// Definir expectativas para as chamadas de mock
	mockUserRepo.EXPECT().GetUsers().Return(users, nil)

	// Criar o UserService
	us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

	// Testar a função GetUsers
	result, err := us.GetUsers()
	assert.Nil(t, err)
	assert.Equal(t, users, result)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	// Defina um usuário que você espera que seja retornado pelo seu repositório
	user := &database.User{
		Name:     "Test",
		Email:    "test@email.com",
		Status:   "Active",
		TenantID: "123",
	}

	// Definir expectativas para as chamadas de mock
	mockUserRepo.EXPECT().GetUser("123").Return(user, nil)

	// Criar o UserService
	us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

	// Testar a função GetUser
	result, err := us.GetUser("123")
	assert.Nil(t, err)
	assert.Equal(t, user, result)
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	// Defina um usuário que você espera que seja retornado pelo seu repositório
	user := &database.User{
		Name:     "Test",
		Email:    "test@email.com",
		Status:   "Active",
		TenantID: "123",
	}

	// Definir expectativas para as chamadas de mock
	mockUserRepo.EXPECT().GetUserByEmail("test@email.com").Return(user, nil)

	// Criar o UserService
	us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

	// Testar a função GetUserByEmail
	result, err := us.GetUserByEmail("test@email.com")
	assert.Nil(t, err)
	assert.Equal(t, user, result)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockIUserRepository(ctrl)
	mockRoleRepo := mock.NewMockIRBCA(ctrl)

	user := &database.UpdateUserParams{
		Name:           "Test User",
		Email:          "test@example.com",
		DocumentType:   sql.NullString{String: "doctype", Valid: true},
		DocumentNumber: sql.NullString{String: "docnumber", Valid: true},
		Phone:          sql.NullString{String: "phone", Valid: true},
	}

	t.Run("should update user successfully", func(t *testing.T) {
		mockUserRepo.EXPECT().UpdateUser(user.ID, user).Return(nil)

		us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

		err := us.UpdateUser(user.ID, user)

		assert.NoError(t, err)
	})

	t.Run("should return error if update fails", func(t *testing.T) {
		mockUserRepo.EXPECT().UpdateUser(user.ID, user).Return(errors.New("update error"))

		us := service.NewServiceUser(mockUserRepo, mockRoleRepo)

		err := us.UpdateUser(user.ID, user)

		assert.Error(t, err)
	})
}
