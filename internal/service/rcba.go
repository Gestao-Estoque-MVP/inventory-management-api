package service

import (
	"log"

	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"nullprogram.com/x/uuid"
)

type RCBAService struct {
	rcba repository.IRBCA
}

func NewRCBAService(rcba repository.IRBCA) *RCBAService {
	return &RCBAService{rcba: rcba}
}

func (r *RCBAService) CreateRoles(role *database.Role) (*database.Role, error) {
	create, err := r.rcba.CreateRoles(&database.Role{
		ID:          uuid.NewGen().NewV4().String(),
		Name:        role.Name,
		Description: role.Description,
	})

	if err != nil {
		log.Printf("Erro na criação dos roles: %v", err)
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) GetRolesPermissions(role string) ([]*database.GetRolesPermissionsRow, error) {
	permissions, err := r.rcba.GetRolesPermissions(role)

	if err != nil {
		log.Printf("Erro em trazer os Roles_Permissions: %v", err)
		return nil, err
	}

	return permissions, nil
}

func (r *RCBAService) CreatePermissions(permission *database.Permission) (*database.Permission, error) {
	create, err := r.rcba.CreatePermissions(&database.Permission{
		ID:          uuid.NewGen().NewV4().String(),
		Name:        permission.Name,
		Description: permission.Description,
	})

	if err != nil {
		log.Printf("Erro ao Criar as permissions: %v", err)
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) CreateRolesPermissions(assignment *database.RolesPermission) (*database.RolesPermission, error) {
	create, err := r.rcba.CreateRolesPermissions(&database.RolesPermission{
		RoleID:       assignment.RoleID,
		PermissionID: assignment.PermissionID,
	})

	if err != nil {
		log.Printf("Erro ao vincular as permissions ao Role: %v", err)
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) GetUsersPermissions(user string) ([]*database.GetUsersPermissionsRow, error) {
	permissions, err := r.rcba.GetUsersPermissions(user)

	if err != nil {
		log.Printf("Erro ao trazer as permissiões do user %v", err)
	}

	return permissions, err
}

func (r *RCBAService) GetRole(name string) (*database.Role, error) {
	get, err := r.rcba.GetRole(name)

	if err != nil {
		log.Printf("Erro ao Trazer os dados role: %v", err)
		return nil, err
	}

	return get, nil
}

func (r *RCBAService) CreateTenant(tenant string) (*database.Tenant, error) {
	create, err := r.rcba.CreateTenant(&database.Tenant{
		ID:   uuid.NewGen().NewV4().String(),
		Name: pgtype.Text{String: tenant, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) GetRoleByID(id string) (*database.Role, error) {
	get, err := r.rcba.GetRoleByID(id)

	if err != nil {
		log.Printf("Erro ao Trazer os dados role by ID %v", err)
		return nil, err
	}

	return get, nil
}
