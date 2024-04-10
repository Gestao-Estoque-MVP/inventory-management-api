package rcba_service

import (
	"log"

	"github.com/diogoX451/inventory-management-api/internal/database"
	rcba_repository "github.com/diogoX451/inventory-management-api/internal/repository/rcba"
)

type IRCBAService interface {
	CreateRoles(role *database.Role) (*database.Role, error)
	GetRolesPermissions(role [16]byte) ([]*database.GetRolesPermissionsRow, error)
	CreatePermissions(permission *database.Permission) (*database.Permission, error)
	CreateRolesPermissions(assignment *database.RolesPermission) (*database.RolesPermission, error)
	CreateUsersPermissions(assignment *database.UsersPermission) (*database.UsersPermission, error)
	CreateUsersRoles(assignment *database.UsersRole) (*database.UsersRole, error)
	GetUsersPermissions(user [16]byte) ([]*database.GetUsersPermissionsRow, error)
	GetRole(name string) (*database.Role, error)
	GetRoleByID(id [16]byte) (*database.Role, error)
}

type RCBAService struct {
	rcba rcba_repository.IRBCA
}

func NewRCBAService(rcba rcba_repository.IRBCA) *RCBAService {
	return &RCBAService{rcba: rcba}
}

func (r *RCBAService) CreateRoles(role *database.Role) (*database.Role, error) {
	create, err := r.rcba.CreateRoles(&database.Role{
		Name:        role.Name,
		Description: role.Description,
	})

	if err != nil {
		log.Printf("Erro na criação dos roles: %v", err)
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) GetRolesPermissions(role [16]byte) ([]*database.GetRolesPermissionsRow, error) {
	permissions, err := r.rcba.GetRolesPermissions(role)

	if err != nil {
		log.Printf("Erro em trazer os Roles_Permissions: %v", err)
		return nil, err
	}

	return permissions, nil
}

func (r *RCBAService) CreatePermissions(permission *database.Permission) (*database.Permission, error) {
	create, err := r.rcba.CreatePermissions(&database.Permission{
		Name:        permission.Name,
		Description: permission.Description,
	})

	if err != nil {
		log.Printf("Erro ao Criar as permissions: %v", err)
		return nil, err
	}

	return create, nil
}

func (r *RCBAService) CreateUsersRoles(assignment *database.UsersRole) (*database.UsersRole, error) {
	create, err := r.rcba.CreateUsersRoles(&database.UsersRole{
		UserID: assignment.UserID,
		RoleID: assignment.RoleID,
	})

	if err != nil {
		log.Printf("Erro ao vincular os roles ao User: %v", err)
		return nil, err
	}

	return create, nil

}

func (r *RCBAService) CreateUsersPermissions(assignment *database.UsersPermission) (*database.UsersPermission, error) {
	create, err := r.rcba.CreateUsersPermissions(&database.UsersPermission{
		UserID:       assignment.UserID,
		PermissionID: assignment.PermissionID,
	})

	if err != nil {
		log.Printf("Erro ao vincular as permissions ao User: %v", err)
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

func (r *RCBAService) GetUsersPermissions(user [16]byte) ([]*database.GetUsersPermissionsRow, error) {
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

func (r *RCBAService) GetRoleByID(id [16]byte) (*database.Role, error) {
	get, err := r.rcba.GetRoleByID(id)

	if err != nil {
		log.Printf("Erro ao Trazer os dados role by ID %v", err)
		return nil, err
	}

	return get, nil
}

func (r *RCBAService) GetRoleUser(id [16]byte) (*database.GetRoleUserRow, error) {
	get, err := r.rcba.GetRoleUser(id)

	if err != nil {
		log.Printf("Error getting role user %s: %v", id, err)
		return nil, err
	}

	return get, nil
}
