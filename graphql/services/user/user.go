package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
)

type userservice struct{}

func NewUserServices() UserService {
	return &userservice{}
}

func (usr *userservice) Register(user model.User) (*model.User, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/register", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return &user, nil
}

func (usr *userservice) Login(user model.User) (*model.User, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/login", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)
	return &user, nil
}

func (usr *userservice) Detail(id *string) (*model.User, error) {
	var user model.User
	requestUrl := fmt.Sprintf("http://localhost:8000/users/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)
	return &user, nil
}
