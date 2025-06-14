// Code generated by MockGen. DO NOT EDIT.
// Source: db/sqlc/agent.go
//
// Generated by this command:
//
//	mockgen -source=db/sqlc/agent.go -destination=db/mock/agent.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	sqlc "github.com/mickaelyoshua/personal_finances/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockAgent is a mock of Agent interface.
type MockAgent struct {
	ctrl     *gomock.Controller
	recorder *MockAgentMockRecorder
	isgomock struct{}
}

// MockAgentMockRecorder is the mock recorder for MockAgent.
type MockAgentMockRecorder struct {
	mock *MockAgent
}

// NewMockAgent creates a new mock instance.
func NewMockAgent(ctrl *gomock.Controller) *MockAgent {
	mock := &MockAgent{ctrl: ctrl}
	mock.recorder = &MockAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgent) EXPECT() *MockAgentMockRecorder {
	return m.recorder
}

// CreateExpense mocks base method.
func (m *MockAgent) CreateExpense(ctx context.Context, arg sqlc.CreateExpenseParams) (sqlc.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpense", ctx, arg)
	ret0, _ := ret[0].(sqlc.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExpense indicates an expected call of CreateExpense.
func (mr *MockAgentMockRecorder) CreateExpense(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpense", reflect.TypeOf((*MockAgent)(nil).CreateExpense), ctx, arg)
}

// CreateIncome mocks base method.
func (m *MockAgent) CreateIncome(ctx context.Context, arg sqlc.CreateIncomeParams) (sqlc.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncome", ctx, arg)
	ret0, _ := ret[0].(sqlc.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncome indicates an expected call of CreateIncome.
func (mr *MockAgentMockRecorder) CreateIncome(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncome", reflect.TypeOf((*MockAgent)(nil).CreateIncome), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockAgent) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAgentMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAgent)(nil).CreateUser), ctx, arg)
}

// DeleteExpense mocks base method.
func (m *MockAgent) DeleteExpense(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpense", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpense indicates an expected call of DeleteExpense.
func (mr *MockAgentMockRecorder) DeleteExpense(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpense", reflect.TypeOf((*MockAgent)(nil).DeleteExpense), ctx, id)
}

// DeleteIncome mocks base method.
func (m *MockAgent) DeleteIncome(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIncome", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIncome indicates an expected call of DeleteIncome.
func (mr *MockAgentMockRecorder) DeleteIncome(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIncome", reflect.TypeOf((*MockAgent)(nil).DeleteIncome), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockAgent) DeleteUser(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockAgentMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockAgent)(nil).DeleteUser), ctx, id)
}

// GetAllExpenses mocks base method.
func (m *MockAgent) GetAllExpenses(ctx context.Context, userID int32) ([]sqlc.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExpenses", ctx, userID)
	ret0, _ := ret[0].([]sqlc.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExpenses indicates an expected call of GetAllExpenses.
func (mr *MockAgentMockRecorder) GetAllExpenses(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExpenses", reflect.TypeOf((*MockAgent)(nil).GetAllExpenses), ctx, userID)
}

// GetAllIncomes mocks base method.
func (m *MockAgent) GetAllIncomes(ctx context.Context, userID int32) ([]sqlc.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIncomes", ctx, userID)
	ret0, _ := ret[0].([]sqlc.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIncomes indicates an expected call of GetAllIncomes.
func (mr *MockAgentMockRecorder) GetAllIncomes(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIncomes", reflect.TypeOf((*MockAgent)(nil).GetAllIncomes), ctx, userID)
}

// GetAllUsers mocks base method.
func (m *MockAgent) GetAllUsers(ctx context.Context) ([]sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockAgentMockRecorder) GetAllUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockAgent)(nil).GetAllUsers), ctx)
}

// GetAllUsersWithDeleted mocks base method.
func (m *MockAgent) GetAllUsersWithDeleted(ctx context.Context) ([]sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersWithDeleted", ctx)
	ret0, _ := ret[0].([]sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersWithDeleted indicates an expected call of GetAllUsersWithDeleted.
func (mr *MockAgentMockRecorder) GetAllUsersWithDeleted(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersWithDeleted", reflect.TypeOf((*MockAgent)(nil).GetAllUsersWithDeleted), ctx)
}

// GetExpense mocks base method.
func (m *MockAgent) GetExpense(ctx context.Context, id int32) (sqlc.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpense", ctx, id)
	ret0, _ := ret[0].(sqlc.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpense indicates an expected call of GetExpense.
func (mr *MockAgentMockRecorder) GetExpense(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpense", reflect.TypeOf((*MockAgent)(nil).GetExpense), ctx, id)
}

// GetIncome mocks base method.
func (m *MockAgent) GetIncome(ctx context.Context, id int32) (sqlc.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncome", ctx, id)
	ret0, _ := ret[0].(sqlc.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncome indicates an expected call of GetIncome.
func (mr *MockAgentMockRecorder) GetIncome(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncome", reflect.TypeOf((*MockAgent)(nil).GetIncome), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockAgent) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockAgentMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockAgent)(nil).GetUserByEmail), ctx, email)
}

// GetUserById mocks base method.
func (m *MockAgent) GetUserById(ctx context.Context, id int32) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockAgentMockRecorder) GetUserById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockAgent)(nil).GetUserById), ctx, id)
}

// HardDeleteUser mocks base method.
func (m *MockAgent) HardDeleteUser(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDeleteUser indicates an expected call of HardDeleteUser.
func (mr *MockAgentMockRecorder) HardDeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDeleteUser", reflect.TypeOf((*MockAgent)(nil).HardDeleteUser), ctx, id)
}

// RestoreUser mocks base method.
func (m *MockAgent) RestoreUser(ctx context.Context, id int32) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreUser", ctx, id)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreUser indicates an expected call of RestoreUser.
func (mr *MockAgentMockRecorder) RestoreUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreUser", reflect.TypeOf((*MockAgent)(nil).RestoreUser), ctx, id)
}

// UpdateExpense mocks base method.
func (m *MockAgent) UpdateExpense(ctx context.Context, arg sqlc.UpdateExpenseParams) (sqlc.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpense", ctx, arg)
	ret0, _ := ret[0].(sqlc.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateExpense indicates an expected call of UpdateExpense.
func (mr *MockAgentMockRecorder) UpdateExpense(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpense", reflect.TypeOf((*MockAgent)(nil).UpdateExpense), ctx, arg)
}

// UpdateIncome mocks base method.
func (m *MockAgent) UpdateIncome(ctx context.Context, arg sqlc.UpdateIncomeParams) (sqlc.Income, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIncome", ctx, arg)
	ret0, _ := ret[0].(sqlc.Income)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIncome indicates an expected call of UpdateIncome.
func (mr *MockAgentMockRecorder) UpdateIncome(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIncome", reflect.TypeOf((*MockAgent)(nil).UpdateIncome), ctx, arg)
}

// UpdateUser mocks base method.
func (m *MockAgent) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, arg)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAgentMockRecorder) UpdateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAgent)(nil).UpdateUser), ctx, arg)
}
