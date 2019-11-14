package sms

import (
	"fmt"
	"wyimserver/common"
	"wyimserver/utils"
)

type SmsServer struct {
	msg SmsModule
}

type SmsModule interface {
	SendCodeSms(map[string]interface{}) //发送验证码短信

}

func (msg *SmsServer) SendCodeSms(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WySmsEndpoint+common.SendCodeSmsUrl, "POST", []byte(GetBodyStr(info)))
}

func GetBodyStr(info map[string]interface{}) (bodyStr string) {
	if tname, ok := info["mobile"]; ok {
		bodyStr += fmt.Sprintf("mobile=%s", tname.(string))
	}

	if deviceId, ok := info["deviceId"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("deviceId=%s", deviceId.(string))
	}

	if templateid, ok := info["templateid"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("templateid=%d", templateid.(int))
	}

	if codeLen, ok := info["codeLen"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("codeLen=%s", codeLen.(string))
	}

	if authCode, ok := info["authCode"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("authCode=%s", authCode.(string))
	}

	if needUp, ok := info["needUp"]; ok {
		bodyStr += fmt.Sprintf("&needUp=%t", needUp.(bool))
	}

	return
}
