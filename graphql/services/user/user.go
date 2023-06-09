package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rulanugrh/graphql/graph/model"
)

type userservice struct{}

func NewUserServices() UserService {
	return &userservice{}
}

func (usr *userservice) Register(user model.User) (*model.Response, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/register", bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (usr *userservice) Login(user model.User) (*model.Response, error) {
	json_data, _ := json.Marshal(&user)
	resp, err := http.NewRequest("POST", "http://localhost:8000/users/login", bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (usr *userservice) Detail(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:8000/users/%d", id)
	resp, err := http.Get(requestUrl)

	var sb strings.Builder
	io.Copy(&sb, resp.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}
