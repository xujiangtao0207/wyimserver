package chatroom

import (
	"fmt"
	"wyimserver/common"
	"wyimserver/utils"
)

type ChatroomServer struct {
	chatroom ChatroomModule
}

type ChatroomModule interface {
	CreateChatroom(map[string]interface{})
	UpdateChatroom(map[string]interface{})
}

func (u *ChatroomServer) CreateChatroom(info map[string]interface{}) ([]byte, int, error) {

	return utils.SendRequestToWy(common.WyImEndpoint+common.CreateChatroomUrl, "POST", []byte(GetBodyStr(info)))
}

func (u *ChatroomServer) UpdateChatroom(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.UpdateChatrommUrl, "POST", []byte(GetBodyStr(info)))
}

func GetBodyStr(info map[string]interface{}) (bodyStr string) {
	//创建
	if creator, ok := info["creator"]; ok {
		bodyStr += fmt.Sprintf("creator=%s", creator.(string))
	}
	//更新
	if roomid, ok := info["roomid"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("roomid=%s", roomid.(string))
	}

	if name, ok := info["name"]; ok {
		bodyStr += fmt.Sprintf("&name=%s", name.(string))
	}
	if announcement, ok := info["announcement"]; ok {
		bodyStr += fmt.Sprintf("&announcement=%s", announcement.(string))
	}
	if ext, ok := info["ext"]; ok {
		bodyStr += fmt.Sprintf("&ext=%s", ext.(string))
	}
	if queuelevel, ok := info["queuelevel"]; ok {
		bodyStr += fmt.Sprintf("&queuelevel=%d", queuelevel.(int))
	}
	//更新
	if needNotify, ok := info["needNotify"]; ok {
		bodyStr += fmt.Sprintf("&needNotify=%s", needNotify.(string))
	}
	if notifyExt, ok := info["notifyExt"]; ok {
		bodyStr += fmt.Sprintf("&notifyExt=%s", notifyExt.(string))
	}
	return bodyStr
}
