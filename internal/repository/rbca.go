package repository

import (
	"context"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRBCA interface {
	CreateRoles(*database.Role) (*database.Role, error)
	GetRole(name string) (*database.Role, error)
	CreatePermissions(*database.Permission) (*database.Permission, error)
	CreateRolesPermissions(*database.RolesPermission) (*database.RolesPermission, error)
	CreateUsersPermissions(*database.UsersPermission) (*database.UsersPermission, error)
	CreateTenant(*database.Tenant) (*database.Tenant, error)
	GetRolesPermissions(role [16]byte) ([]*database.GetRolesPermissionsRow, error)
	GetUsersPermissions(user [16]byte) ([]*database.GetUsersPermissionsRow, error)
	GetRoleByID(id [16]byte) (*database.Role, error)
	GetRoleUser(id [16]byte) (*database.GetRoleUserRow, error)
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

func (r *RBCARepository) GetRolesPermissions(role [16]byte) ([]*database.GetRolesPermissionsRow, error) {
	getRole, err := r.DB.GetRolesPermissions(context.Background(), pgtype.UUID{Bytes: role, Valid: true})

	if err != nil {
		return nil, err
	}

	pointers := make([]*database.GetRolesPermissionsRow, len(getRole))
	for i := range getRole {
		pointers[i] = &getRole[i]
	}

	return pointers, nil
}

func (r *RBCARepository) GetUsersPermissions(user [16]byte) ([]*database.GetUsersPermissionsRow, error) {
	getUser, err := r.DB.GetUsersPermissions(context.Background(), pgtype.UUID{Bytes: user, Valid: true})

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

func (r *RBCARepository) GetRoleByID(id [16]byte) (*database.Role, error) {
	get, err := r.DB.GetRoleByID(context.Background(), pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		return nil, err
	}

	return &get, nil
}

func (r *RBCARepository) GetRoleUser(id [16]byte) (*database.GetRoleUserRow, error) {
	get, err := r.DB.GetRoleUser(context.Background(), pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		return nil, err
	}

	return &get, nil
}
