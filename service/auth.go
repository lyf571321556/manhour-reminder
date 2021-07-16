package service

import (
	"encoding/json"
	"fmt"
	"github.com/lyf571321556/qiye-wechat-bot-api/api"
	"io/ioutil"
	"net/http"
)

var (
	AUTH_LOGIN = "/auth/login"
	ITEMS_GQL  = "/team/%s/items/graphql"
)

type AuthInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   string `json:"uuid"`
	Token    string `json:"token,omitempty"`
}

// Login for token /*
func Login(url, username, password string) (user AuthInfo, err error) {
	result := AuthInfo{
		Email:    username,
		Password: password,
	}
	var userJson []byte
	if userJson, err = json.Marshal(result); err != nil {
		return user, err
	}
	request, err := api.NewRequest(http.MethodPost, url, userJson)
	if err != nil {
		return user, err
	}

	var loginInfo = new(struct {
		UserInfo AuthInfo `json:"user"`
	})
	_, err = api.ExecuteHTTP(request, func(resp *http.Response) error {
		rawResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error:%s", string(rawResp))
		}
		if err = json.Unmarshal(rawResp, loginInfo); err != nil {
			return fmt.Errorf("unknown response: %w\nraw response: %s", err, rawResp)
		}
		return nil
	})
	if err != nil {
		return user, err
	}

	return loginInfo.UserInfo, nil
}
