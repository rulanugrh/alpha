package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rulanugrh/alpha/apigateway/entity/user"
)

type userservice struct{}

func NewUserServices() UserService {
	return &userservice{}
}

func (usr *userservice) Register(user user.UserModel) (http.Response, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/register", bytes.NewBuffer(json_data))
	if err != nil {
		return http.Response{
			Status:     "cannot register",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (usr *userservice) Login(user user.UserModel) (http.Response, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/login", bytes.NewBuffer(json_data))
	if err != nil {
		return http.Response{
			Status:     "cannot login",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (usr *userservice) Detail(id uint) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:8000/users/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return http.Response{
			Status:     "cannot register",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp, nil
}
