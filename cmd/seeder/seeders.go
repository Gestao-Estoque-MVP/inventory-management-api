package main

import (
	"context"
	"log"
	"os"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func getUUID() pgtype.UUID {
	id, _ := uuid.NewV4()
	return pgtype.UUID{Bytes: id, Valid: true}
}

var roles = []database.Role{
	{
		ID:          getUUID(),
		Name:        "superuser",
		Description: "Papeis dedicado aos users",
	},
	{
		ID:          getUUID(),
		Name:        "admin",
		Description: "Papeis dedicado aos admin",
	},
	{
		ID:          getUUID(),
		Name:        "user",
		Description: "Papeis dedicado aos user",
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
	{
		ID:          getUUID(),
		Name:        "read",
		Description: "Papeis dedicado aos read",
	},
	{
		ID:          getUUID(),
		Name:        "list",
		Description: "Papeis dedicado aos list",
	},
	{
		ID:          getUUID(),
		Name:        "view",
		Description: "Papeis dedicado aos view",
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

	return nil
}

func init() {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

}

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
	}

	defer func() {
		db.Close()
	}()
	queries := database.New(db)
	err = Seed(queries)
	if err != nil {
		log.Printf("Error creating database", err)
	}
}
