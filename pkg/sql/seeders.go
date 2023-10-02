package sql

import (
	"context"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/database"
	token "github.com/diogoX451/inventory-management-api/pkg/Token"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func getUUID() pgtype.UUID {
	id, _ := uuid.NewV4()
	return pgtype.UUID{Bytes: id, Valid: true}
}

var roles = []database.Role{
	{
		ID:          getUUID(),
		Name:        "users",
		Description: "Papeis dedicado aos users",
	},
	{
		ID:          getUUID(),
		Name:        "admin",
		Description: "Papeis dedicado aos admin",
	},
}

var permissions = []database.Permission{
	{
		ID:          getUUID(),
		Name:        "create",
		Description: "Papeis dedicado aos create",
	},
	{
		ID:          getUUID(),
		Name:        "delete",
		Description: "Papeis dedicado aos delete",
	},
	{
		ID:          getUUID(),
		Name:        "update",
		Description: "Papeis dedicado aos update",
	},
}

var rolePermissions = []database.RolesPermission{
	{
		ID:           getUUID(),
		RoleID:       roles[0].ID,
		PermissionID: permissions[0].ID,
	},
	{
		ID:           getUUID(),
		RoleID:       roles[1].ID,
		PermissionID: permissions[1].ID,
	},
	{
		ID:           getUUID(),
		RoleID:       roles[0].ID,
		PermissionID: permissions[1].ID,
	},
	{
		ID:           getUUID(),
		RoleID:       roles[0].ID,
		PermissionID: permissions[2].ID,
	},
	{
		ID:           getUUID(),
		RoleID:       roles[1].ID,
		PermissionID: permissions[0].ID,
	},
	{
		ID:           getUUID(),
		RoleID:       roles[1].ID,
		PermissionID: permissions[2].ID,
	},
}

var company = []database.Tenant{
	{
		ID:        getUUID(),
		Name:      pgtype.Text{String: "SwiftStock", Valid: true},
		TaxCode:   pgtype.Text{String: "00000000001", Valid: true},
		Type:      database.NullTenantType{TenantType: database.TenantTypeCustomer, Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now().Local()},
	},
}

var users = []database.User{
	{
		ID:             getUUID(),
		Name:           pgtype.Text{String: "Admin Teste", Valid: true},
		Email:          "admin@teste.com",
		DocumentType:   pgtype.Text{String: "CPF", Valid: true},
		DocumentNumber: pgtype.Text{String: "00010110203", Valid: true},
		Password:       pgtype.Text{String: token.HashPassword("123456"), Valid: true},
		Status:         database.UserStatusActive,
		TenantID:       company[0].ID,
		CreatedAt:      pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
	},
}

var usersRoles = []database.UsersRole{
	{
		ID:     getUUID(),
		UserID: users[0].ID,
		RoleID: roles[1].ID,
	},
}

func Seed(db *database.Queries) error {
	for _, role := range roles {
		createRoleParams := &database.CreateRoleParams{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		}
		_, err := db.CreateRole(context.Background(), *createRoleParams)
		if err != nil {
			return err
		}
	}

	for _, permission := range permissions {
		createPermissionsParams := &database.CreatePermissionsParams{
			ID:          permission.ID,
			Name:        permission.Name,
			Description: permission.Description,
		}
		_, err := db.CreatePermissions(context.Background(), *createPermissionsParams)
		if err != nil {
			return err
		}
	}

	for _, rp := range rolePermissions {
		_, err := db.CreateRolePermissions(context.Background(), database.CreateRolePermissionsParams{
			ID:           rp.ID,
			RoleID:       rp.RoleID,
			PermissionID: rp.PermissionID,
		})
		if err != nil {
			return err
		}
	}

	for _, ct := range company {
		_, err := db.CreateTenant(context.Background(), database.CreateTenantParams{
			ID:        ct.ID,
			Name:      ct.Name,
			TaxCode:   ct.TaxCode,
			Type:      ct.Type,
			CreatedAt: ct.CreatedAt,
		})
		if err != nil {
			return err
		}
	}

	for _, cu := range users {
		tk, _ := token.GeneratedToken()

		_, err := db.CreatePreRegisterUser(context.Background(), database.CreatePreRegisterUserParams{
			ID:            cu.ID,
			Name:          cu.Name,
			Email:         cu.Email,
			Status:        cu.Status,
			ID_2:          getUUID(),
			Type:          database.TypeNumberHome,
			Number:        "62999722708",
			IsPrimary:     true,
			CreatedAt_2:   pgtype.Timestamp{Time: time.Now().Local(), Valid: true},
			RegisterToken: pgtype.Text{String: "", Valid: true},
			TenantID:      cu.TenantID,
			CreatedAt:     cu.CreatedAt,
		})

		if err != nil {
			return err
		}

		_, err = db.CompleteRegisterUser(context.Background(), database.CompleteRegisterUserParams{
			DocumentType:   cu.DocumentType,
			DocumentNumber: cu.DocumentNumber,
			Password:       cu.Password,
			Status:         database.UserStatusActive,
			RegisterToken:  pgtype.Text{String: tk, Valid: true},
		})

		if err != nil {
			return err
		}

	}

	for _, ur := range usersRoles {
		_, err := db.CreateUsersRoles(context.Background(), database.CreateUsersRolesParams{
			ID:     ur.ID,
			UserID: ur.ID,
			RoleID: ur.RoleID,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
