package msg

import (
	"fmt"
	"wyimserver/common"
	"wyimserver/utils"
)

type MsgServer struct {
	msg MsgModule
}

type MsgModule interface {
	SendMsg(map[string]interface{})            //发送普通消息
	SendAttachMsg(map[string]interface{})      //发送自定义系统通知
	SendBatchAttachMsg(map[string]interface{}) //批量发送点对点自定义系统通知

}

func (msg *MsgServer) SendMsg(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.SendMsgUrl, "POST", []byte(GetBodyStr(info)))
}
func (msg *MsgServer) SendAttachMsg(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.SendAttachMsgUrl, "POST", []byte(GetBodyStr(info)))
}

func (msg *MsgServer) SendBatchAttachMsg(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.SendBatchAttachMsgUrl, "POST", []byte(GetBodyStr(info)))
}

func GetBodyStr(info map[string]interface{}) (bodyStr string) {
	if tname, ok := info["from"]; ok {
		bodyStr += fmt.Sprintf("from=%s", tname.(string))
	}

	if ope, ok := info["ope"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("ope=%d", ope.(int))
	}

	if types, ok := info["type"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("type=%d", types.(int))
	}

	if body, ok := info["body"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("body=%s", body.(string))
	}

	if fromAccid, ok := info["fromAccid"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("fromAccid=%s", fromAccid.(string))
	}

	if toAccids, ok := info["toAccids"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("toAccids=%s", toAccids.(string))
	}

	if msgtype, ok := info["msgtype"]; ok {
		bodyStr += fmt.Sprintf("&msgtype=%d", msgtype.(int))
	}

	if to, ok := info["to"]; ok {
		bodyStr += fmt.Sprintf("&to=%s", to.(string))
	}

	if attach, ok := info["attach"]; ok {
		bodyStr += fmt.Sprintf("&attach=%s", attach.(string))
	}
	if pushcontent, ok := info["pushcontent"]; ok {
		bodyStr += fmt.Sprintf("&pushcontent=%s", pushcontent.(string))
	}

	if payload, ok := info["payload"]; ok {
		bodyStr += fmt.Sprintf("&payload=%s", payload.(string))
	}

	if sound, ok := info["sound"]; ok {
		bodyStr += fmt.Sprintf("&sound=%s", sound.(string))
	}

	if save, ok := info["save"]; ok {
		bodyStr += fmt.Sprintf("&save=%d", save.(int))
	}

	if option, ok := info["option"]; ok {
		bodyStr += fmt.Sprintf("&option=%s", option.(string))
	}

	if antispam, ok := info["antispam"]; ok {
		bodyStr += fmt.Sprintf("&antispam=%s", antispam.(string))
	}

	if antispamCustom, ok := info["antispamCustom"]; ok {
		bodyStr += fmt.Sprintf("&antispamCustom=%s", antispamCustom.(string))
	}

	return
}
