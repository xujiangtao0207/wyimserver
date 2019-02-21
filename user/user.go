package user

import (
	"fmt"
	"wyimserver/common"
	"wyimserver/utils"
)

type UserServer struct {
	user UserModule
}

type UserModule interface {
	CreateUser(map[string]interface{})
	UpdateUser(map[string]interface{})
}

func (u *UserServer) CreateUser(info map[string]interface{}) ([]byte, int, error) {

	return utils.SendRequestToWy(common.WyImEndpoint+common.CreateUserUrl, "POST", []byte(GetBodyStr(info)))
}

func (u *UserServer) UpdateUser(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.UpdateUserUrl, "POST", []byte(GetBodyStr(info)))
}

func GetBodyStr(info map[string]interface{}) (bodyStr string) {
	if accid, ok := info["accid"]; ok {
		bodyStr += fmt.Sprintf("accid=%s", accid.(string))
	}

	if name, ok := info["name"]; ok {
		bodyStr += fmt.Sprintf("&name=%s", name.(string))
	}

	if props, ok := info["props"]; ok {
		bodyStr += fmt.Sprintf("&props=%s", props.(string))
	}
	if icon, ok := info["icon"]; ok {
		bodyStr += fmt.Sprintf("&icon=%s", icon.(string))
	}
	if token, ok := info["token"]; ok {
		bodyStr += fmt.Sprintf("&token=%s", token.(string))
	}

	if sign, ok := info["sign"]; ok {
		bodyStr += fmt.Sprintf("&sign=%s", sign.(string))
	}
	if email, ok := info["email"]; ok {
		bodyStr += fmt.Sprintf("&email=%s", email.(string))
	}
	if birth, ok := info["birth"]; ok {
		bodyStr += fmt.Sprintf("&birth=%s", birth.(string))
	}
	if mobile, ok := info["mobile"]; ok {
		bodyStr += fmt.Sprintf("&mobile=%s", mobile.(string))
	}
	if gender, ok := info["gender"]; ok {
		bodyStr += fmt.Sprintf("&gender=%d", gender.(int))
	}
	if ex, ok := info["ex"]; ok {
		bodyStr += fmt.Sprintf("&ex=%s", ex.(string))
	}
	return
}
