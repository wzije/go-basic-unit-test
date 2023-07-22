package http_server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var UserData = []User{
	{ID: "1", Name: "Mr. One"},
	{ID: "2", Name: "Mr. Two"},
	{ID: "3", Name: "Mr. Three"},
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	Get() (*[]User, error)
	FindByID(id string) (*User, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) Get() (*[]User, error) {
	return &UserData, nil
}

func (u *userService) FindByID(id string) (*User, error) {
	index := -1

	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("data not found")
	}

	return &UserData[index], nil
}

func GetUsers(service UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		users, err := service.Get()

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		userJson, _ := json.Marshal(users)

		_, err = writer.Write(userJson)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}
}

func GetUser(service UserService) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		id := strings.TrimPrefix(request.URL.Path, "/users/")

		if id == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := service.FindByID(id)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		userJson, _ := json.Marshal(user)

		_, err = writer.Write(userJson)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}
}
