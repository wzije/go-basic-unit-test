package http_server

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Get() (*[]User, error) {
	args := m.Called()
	return args.Get(0).(*[]User), args.Error(1)
}

func (m *MockService) FindByID(id string) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("error")
	} else {
		return args.Get(0).(*User), nil
	}
}

func TestUserService_FindByID(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		id := "1"
		user := UserData[0]
		userMock := new(MockService)
		userMock.On("FindById", id).Return(&user, nil)

		userService := NewUserService()
		u, err := userService.FindByID(id)

		assert.Nil(t, err)
		assert.Equal(t, user.ID, u.ID)
	})

	t.Run("should error", func(t *testing.T) {
		id := "10"
		userMock := new(MockService)
		userMock.On("FindById", id).Return(nil, errors.New("data not found"))

		userService := NewUserService()
		_, err := userService.FindByID(id)

		assert.NotNil(t, err)
	})
}

func TestUserService_Get(t *testing.T) {
	t.Run(fmt.Sprintf("should success and total data should %d", len(UserData)), func(t *testing.T) {
		users := UserData
		userMock := new(MockService)
		userMock.On("FindById", mock.Anything).Return(&users, nil)

		userService := NewUserService()
		result, err := userService.Get()

		assert.Nil(t, err)
		assert.Equal(t, len(UserData), len(*result))
	})
}

func TestHandler_GetUser(t *testing.T) {

	t.Run("get user by available id should success", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)

		userMock := new(MockService)
		user := User{ID: "1", Name: "Mr. One"}

		userMock.On("FindByID", mock.Anything).Return(&user, nil)
		recorder := httptest.NewRecorder()

		handler := GetUser(userMock)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		expected := `{"id": "1","name": "Mr. One"}`
		assert.JSONEq(t, expected, recorder.Body.String())
	})

	t.Run("get user by without id should bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/", nil)

		userMock := new(MockService)
		userMock.On("FindByID", mock.Anything).Return(nil, err)

		recorder := httptest.NewRecorder()

		handler := GetUser(userMock)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("get user by unavailable id should bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/10", nil)

		userMock := new(MockService)
		userMock.On("FindByID", mock.Anything).Return(nil, err)
		recorder := httptest.NewRecorder()

		handler := GetUser(userMock)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

}

func TestHandler_GetUsers(t *testing.T) {
	t.Run("get users should success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)

		userMock := new(MockService)
		userMock.On("Get", mock.Anything).Return(&[]User{}, nil)
		recorder := httptest.NewRecorder()

		handler := GetUsers(userMock)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
