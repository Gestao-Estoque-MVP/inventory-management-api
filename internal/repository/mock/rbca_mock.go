// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/rbca.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	database "github.com/diogoX451/inventory-management-api/internal/database"
	gomock "github.com/golang/mock/gomock"
)

// MockIRBCA is a mock of IRBCA interface.
type MockIRBCA struct {
	ctrl     *gomock.Controller
	recorder *MockIRBCAMockRecorder
}

// MockIRBCAMockRecorder is the mock recorder for MockIRBCA.
type MockIRBCAMockRecorder struct {
	mock *MockIRBCA
}

// NewMockIRBCA creates a new mock instance.
func NewMockIRBCA(ctrl *gomock.Controller) *MockIRBCA {
	mock := &MockIRBCA{ctrl: ctrl}
	mock.recorder = &MockIRBCAMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRBCA) EXPECT() *MockIRBCAMockRecorder {
	return m.recorder
}

// CreatePermissions mocks base method.
func (m *MockIRBCA) CreatePermissions(arg0 *database.Permission) (*database.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermissions", arg0)
	ret0, _ := ret[0].(*database.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermissions indicates an expected call of CreatePermissions.
func (mr *MockIRBCAMockRecorder) CreatePermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermissions", reflect.TypeOf((*MockIRBCA)(nil).CreatePermissions), arg0)
}

// CreateRoles mocks base method.
func (m *MockIRBCA) CreateRoles(arg0 *database.Role) (*database.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoles", arg0)
	ret0, _ := ret[0].(*database.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoles indicates an expected call of CreateRoles.
func (mr *MockIRBCAMockRecorder) CreateRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoles", reflect.TypeOf((*MockIRBCA)(nil).CreateRoles), arg0)
}

// CreateRolesPermissions mocks base method.
func (m *MockIRBCA) CreateRolesPermissions(arg0 *database.RolesPermission) (*database.RolesPermission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRolesPermissions", arg0)
	ret0, _ := ret[0].(*database.RolesPermission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRolesPermissions indicates an expected call of CreateRolesPermissions.
func (mr *MockIRBCAMockRecorder) CreateRolesPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRolesPermissions", reflect.TypeOf((*MockIRBCA)(nil).CreateRolesPermissions), arg0)
}

// CreateTenant mocks base method.
func (m *MockIRBCA) CreateTenant(arg0 *database.Tenant) (*database.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", arg0)
	ret0, _ := ret[0].(*database.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant.
func (mr *MockIRBCAMockRecorder) CreateTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockIRBCA)(nil).CreateTenant), arg0)
}

// CreateUsersPermissions mocks base method.
func (m *MockIRBCA) CreateUsersPermissions(arg0 *database.UsersPermission) (*database.UsersPermission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsersPermissions", arg0)
	ret0, _ := ret[0].(*database.UsersPermission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUsersPermissions indicates an expected call of CreateUsersPermissions.
func (mr *MockIRBCAMockRecorder) CreateUsersPermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsersPermissions", reflect.TypeOf((*MockIRBCA)(nil).CreateUsersPermissions), arg0)
}

// GetRole mocks base method.
func (m *MockIRBCA) GetRole(name string) (*database.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", name)
	ret0, _ := ret[0].(*database.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockIRBCAMockRecorder) GetRole(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIRBCA)(nil).GetRole), name)
}

// GetRolesPermissions mocks base method.
func (m *MockIRBCA) GetRolesPermissions(role string) ([]*database.GetRolesPermissionsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRolesPermissions", role)
	ret0, _ := ret[0].([]*database.GetRolesPermissionsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRolesPermissions indicates an expected call of GetRolesPermissions.
func (mr *MockIRBCAMockRecorder) GetRolesPermissions(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRolesPermissions", reflect.TypeOf((*MockIRBCA)(nil).GetRolesPermissions), role)
}

// GetUsersPermissions mocks base method.
func (m *MockIRBCA) GetUsersPermissions(user string) ([]*database.GetUsersPermissionsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersPermissions", user)
	ret0, _ := ret[0].([]*database.GetUsersPermissionsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersPermissions indicates an expected call of GetUsersPermissions.
func (mr *MockIRBCAMockRecorder) GetUsersPermissions(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersPermissions", reflect.TypeOf((*MockIRBCA)(nil).GetUsersPermissions), user)
}