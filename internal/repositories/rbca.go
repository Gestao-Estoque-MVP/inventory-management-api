package repositories

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type IRBCA interface {
	CreateRoles(*database.Role) (*database.Role, error)
	CreatePermissions(*database.Permission) (*database.Permission, error)
	CreateRolesPermissions(*database.RolesPermission) (*database.RolesPermission, error)
	GetRole(*database.GetRoleRow) (*database.GetRoleRow, error)
}

type RBCARepository struct {
	DB *database.Queries
}

func NewRBCARepository(DB *database.Queries) *RBCARepository {
	return &RBCARepository{
		DB: DB,
	}
}

func (r *RBCARepository) CreateRole(role *database.Role) (*database.Role, error) {
	createRole, err := r.DB.CreateRole(context.Background(), database.CreateRoleParams{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	})

	if err != nil {
		return nil, err
	}

	return &createRole, nil
}

func (r *RBCARepository) CreatePermissions(permissions *database.Permission) (*database.Permission, error) {
	create, err := r.DB.CreatePermissions(context.Background(), database.CreatePermissionsParams{
		ID:          permissions.ID,
		Name:        permissions.Name,
		Description: permissions.Description,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (r *RBCARepository) CreateRolesPermissions(rolesPermissions *database.RolesPermission) (*database.RolesPermission, error) {
	create, err := r.DB.CreateRolePermissions(context.Background(), database.CreateRolePermissionsParams{
		RoleID:       rolesPermissions.RoleID,
		PermissionID: rolesPermissions.RoleID,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (r *RBCARepository) GetRole(role string) (*database.GetRoleRow, error) {
	getRole, err := r.DB.GetRole(context.Background(), role)

	if err != nil {
		return nil, err
	}

	return &getRole, nil
}
