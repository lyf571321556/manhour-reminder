package service

import (
	"encoding/json"
	"fmt"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"net/http"
)

var (
	AUTH_LOGIN = "/auth/login"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

// Login for token /*
func Login(url, username, password string) (token string, err error) {
	user := User{
		Email:    username,
		Password: password,
	}
	var userJson []byte
	if userJson, err = json.Marshal(user); err != nil {
		return token, err
	}
	println(string(userJson))
	request, err := api.NewRequest(http.MethodPost, url, userJson)
	if err != nil {
		return token, err
	}

	var loginInfo = new(struct {
		UserInfo User `json:"user"`
	})
	_, err = api.ExecuteHTTP(request, func(rawResp []byte) error {
		if err = json.Unmarshal(rawResp, loginInfo); err != nil {
			return fmt.Errorf("unknown response: %w\nraw response: %s", err, rawResp)
		}
		return nil
	})
	if err != nil {
		return token, err
	}

	return loginInfo.UserInfo.Token, nil
}
