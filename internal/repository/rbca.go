package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
)

type IRBCA interface {
	CreateRoles(*database.Role) (*database.Role, error)
	GetRole(name string) (*database.Role, error)
	CreatePermissions(*database.Permission) (*database.Permission, error)
	CreateRolesPermissions(*database.RolesPermission) (*database.RolesPermission, error)
	CreateUsersPermissions(*database.UsersPermission) (*database.UsersPermission, error)
	CreateTenant(*database.Tenant) (*database.Tenant, error)
	GetRolesPermissions(role string) ([]*database.GetRolesPermissionsRow, error)
	GetUsersPermissions(user string) ([]*database.GetUsersPermissionsRow, error)
}

type RBCARepository struct {
	DB *database.Queries
}

func NewRBCARepository(DB *database.Queries) *RBCARepository {
	return &RBCARepository{
		DB: DB,
	}
}

func (r *RBCARepository) CreateRoles(role *database.Role) (*database.Role, error) {
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

func (r *RBCARepository) GetRole(name string) (*database.Role, error) {
	get, err := r.DB.GetRole(context.Background(), name)

	if err != nil {
		return nil, err
	}

	return &get, nil
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
		PermissionID: rolesPermissions.PermissionID,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (r *RBCARepository) CreateUsersPermissions(usersPermissions *database.UsersPermission) (*database.UsersPermission, error) {
	create, err := r.DB.CreateUsersPermissions(context.Background(), database.CreateUsersPermissionsParams{
		UserID:       usersPermissions.UserID,
		PermissionID: usersPermissions.PermissionID,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (r *RBCARepository) GetRolesPermissions(role string) ([]*database.GetRolesPermissionsRow, error) {
	getRole, err := r.DB.GetRolesPermissions(context.Background(), role)

	if err != nil {
		return nil, err
	}

	pointers := make([]*database.GetRolesPermissionsRow, len(getRole))
	for i := range getRole {
		pointers[i] = &getRole[i]
	}

	return pointers, nil
}

func (r *RBCARepository) GetUsersPermissions(user string) ([]*database.GetUsersPermissionsRow, error) {
	getUser, err := r.DB.GetUsersPermissions(context.Background(), user)

	if err != nil {
		return nil, err
	}

	pointer := make([]*database.GetUsersPermissionsRow, len(getUser))
	for i := range getUser {
		pointer[i] = &getUser[i]
	}

	return pointer, nil
}

func (r *RBCARepository) CreateTenant(tenant *database.Tenant) (*database.Tenant, error) {
	create, err := r.DB.CreateTenant(context.Background(), database.CreateTenantParams{
		ID:   tenant.ID,
		Name: tenant.Name,
	})

	if err != nil {
		return nil, err
	}

	return &create, nil
}
