package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"wyimserver/chatroom"
	"wyimserver/msg"
	"wyimserver/team"
	"wyimserver/user"
)

func main() {
	// UpdateUserUinfo()
	// GetUinfos()
	// TestGetJoinTeamsForUser()
	// TestSendMsg()
	// TestSendAttachMsg()
	// TestCreateTeam()
	// TestCreateUser()
	// TestAddUserToTeam()

	// RemoveTeam()
	// ClearGroupAll()
	// UpdateTeamGroupInfo()
	// TestGetTeamDetail("2560383622")
	// TestKickUserToTeam()
	KickAllUserToTeam()
}

type JoinTeamsForUserResp struct {
	Code  int     `json:"code"`
	Count int     `json:"count"`
	Infos []Infos `json:"infos"`
}
type Infos struct {
	Owner    string `json:"owner"`
	Tname    string `json:"tname"`
	Maxusers int    `json:"maxusers"`
	Tid      int    `json:"tid"`
	Size     int    `json:"size"`
	Custom   string `json:"custom"`
}

func ClearGroupAll() {
	// var uids = []int{101272265, 101371394, 101307239, 100038367, 101374697, 101387045, 100002033, 101272266,
	// 101272262, 101272263, 101272264, 101272267, 101272260, 100001356, 101273415, 101307226,100002398}
	var uids = []int{100000001}
	for i := 0; i < len(uids); i++ {
		ClearGroup(fmt.Sprintf("%d", uids[i]))
	}
}

func ClearGroup(accid string) {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["accid"] = accid
	bytes, code, err := teams.GetJoinTeamsForUser(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
		return
	}

	var resp JoinTeamsForUserResp
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		logs.Error("解析用户所在群列表信息异常,错误信息:%v,请求响应数据[%s]", err, string(bytes))
		return
	}

	for i := 0; i < len(resp.Infos); i++ {
		var info = resp.Infos[i]
		if info.Owner != accid {
			continue
		}
		RemoveTeamByOwnerAndTid(info.Owner, fmt.Sprintf("%d", info.Tid))
	}
}

func TestGetJoinTeamsForUser() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["accid"] = fmt.Sprintf("%d", 91004904)
	bytes, code, err := teams.GetJoinTeamsForUser(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestSendMsg() {
	var msgs = new(msg.MsgServer)
	var info = make(map[string]interface{})
	info["from"] = "1234aaabaaa"
	info["ope"] = 0
	info["type"] = 0
	info["to"] = "101272265"
	info["body"] = `{
		"msg":"你是猪",
		"type":101,
		"from":"橙瓜",
		"content":"橙瓜邀请你加入",
		"nickname":"https://aa.com"
		}`
	bytes, code, err := msgs.SendMsg(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestSendAttachMsg() {
	var msgs = new(msg.MsgServer)
	var info = make(map[string]interface{})
	info["from"] = "1234aaabaaa"
	info["msgtype"] = 0
	info["to"] = "101272265"
	info["attach"] = `{
		"msg":"倪书洲是你大爷",
		"type":101,
		"from":"橙瓜",
		"content":"橙瓜邀请你加入",
		"nickname":"https://aa.com"
		}`
	info["save"] = 1
	bytes, code, err := msgs.SendAttachMsg(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}

}

//查询群详情
type GetTeamResp struct {
	Code  int   `json:"code"`
	Tinfo Tinfo `json:"tinfo"`
}

type Tinfo struct {
	Tid          int      `json:"tid"`
	Owner        Owner    `json:"owner"`
	Members      []Member `json:"members"`
	Custom       string   `json:"custom"`
	ClientCustom string   `json:"clientCustom"`
}

type Owner struct {
	Accid    string `json:"accid"`
	Nickname string `json:"nick"`
	Custom   string `json:"custom"`
}

type Member struct {
	Nickname string `json:"nick"`
	Accid    string `json:"accid"`
	Custom   string `json:"custom"`
}

var ErrorsNoRows = errors.New("不存在")

func TestGetTeamDetail(tid string) ([]Member, Owner, string, string, error) {
	var teams = new(team.TeamServer)
	var owner Owner
	var info = make(map[string]interface{})
	info["tid"] = tid

	bytes, code, err := teams.GetTeamDetail(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
	var resp GetTeamResp
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		logs.Error("解析群详情异常,错误信息:%v", err)
		return nil, owner, "", "", err
	}

	if resp.Code == 200 {
		return resp.Tinfo.Members, resp.Tinfo.Owner, resp.Tinfo.Custom, resp.Tinfo.ClientCustom, nil
	} else {
		if resp.Code == 414 {
			return nil, owner, "", "", ErrorsNoRows
		}
	}
	return nil, owner, "", "", nil
}

func KickAllUserToTeam() {
	var tids = []string{"2560393181", "2560387530", "2560377941", "2560381653", "2560383622", "2578282280"}
	// var tids = []string{"2560383622"}
	for j := 0; j < len(tids); j++ {

		var tid = tids[j]
		members, owner, _, _, _ := TestGetTeamDetail(tid)

		for i := 0; i < len(members); i++ {
			var member = members[i]
			uid, _ := strconv.Atoi(member.Accid)
			if uid < 100000000 && uid > 0 {
				KickUserToTeam(tid, owner.Accid, fmt.Sprintf("[\"%s\"]", member.Accid))
			}
		}
	}

}

func KickUserToTeam(tid, owner, members string) {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = tid
	info["owner"] = owner
	info["members"] = members
	info["msg"] = "踢人出群"
	info["magree"] = 0
	info["joinmode"] = 0
	logs.Info("打印踢出请求体[%v]", info)
	bytes, code, err := teams.KickUserToTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func TestCreateTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["owner"] = "100000001"
	info["tname"] = "拼字公共房"
	info["members"] = "[\"100000001\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 1
	info["joinmode"] = 0
	info["teamMemberLimit"] = 20
	info["custom"] = "{\"room_uuid\": \"5dd251aeec8f413d95a7a786d75744e0\"}"

	// info["custom"] = "12345"
	bytes, code, err := teams.CreateTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

var teams = new(team.TeamServer)

func RemoveTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["owner"] = "101304569"
	//101374698
	info["tid"] = "2560328383"
	// 2559298425
	bytes, code, err := teams.RemoveTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func RemoveTeamByOwnerAndTid(owner, Tid string) {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["owner"] = owner
	//101374698
	info["tid"] = Tid
	bytes, code, err := teams.RemoveTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestAddUserToTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = 1621410526
	info["owner"] = "101399999"
	info["members"] = "[\"101304574\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 0
	info["joinmode"] = 0
	info["attach"] = "-------"
	bytes, code, err := teams.AddUserToTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func TestKickUserToTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = "2560383622"
	info["owner"] = "100000001"
	info["members"] = "[\"91012472\",\"91012469\",\"91012468\"]"
	info["msg"] = "踢人出群"
	info["magree"] = 0
	info["joinmode"] = 0
	bytes, code, err := teams.KickUserToTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

// func TestGetTeamDetail() {
// 	var teams = new(team.TeamServer)
// 	var info = make(map[string]interface{})
// 	info["tid"] = 1602192891
// 	info["owner"] = "1234aaab"
// 	info["members"] = "[\"1234aaabaaa\"]"
// 	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
// 	info["magree"] = 0
// 	info["joinmode"] = 0
// 	bytes, code, err := teams.GetTeamDetail(info)
// 	logs.Info("请求响应消息[%s]", string(bytes))
// 	logs.Info("请求响应消息[%d]", code)
// 	if err != nil {
// 		logs.Error("创建用户异常,错误信息:%v", err)
// 	}
// }

func TestUpdateTeamNick() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = 1602192891
	info["owner"] = "1234aaab"
	info["accid"] = "1234aaab"

	info["nick"] = "橙瓜码字拼字群"
	bytes, code, err := teams.UpdateTeamNick(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestChangeTeamOwner() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = 1602192891
	info["owner"] = "1234aaabaaa"
	info["newowner"] = "1234aaab"

	info["leave"] = 1
	bytes, code, err := teams.ChangeTeamOwner(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestUpdateUser() {
	var users = new(user.UserServer)
	var info = make(map[string]interface{})
	info["accid"] = "1234aaab"
	info["name"] = "xujiangtao3"
	bytes, code, err := users.UpdateUser(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func TestUpdateUserUinfo() {
	var users = new(user.UserServer)
	var info = make(map[string]interface{})
	info["accid"] = "101272265"
	info["name"] = "xujiangtao3"
	bytes, code, err := users.UpdateUser(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestCreateUser() {
	var users = new(user.UserServer)
	var info = make(map[string]interface{})
	info["accid"] = "101272265"
	info["name"] = "1234"
	info["ex"] = "http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJzogx0ozXBPZj42iapodEK1FfodpSqaKZmynfxyd5nibln44HgUb0Tlc3PZT8qEcLy5VIG2wRoviaxQ/132"
	info["props"] = "{\"avatar\":\"http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJzogx0ozXBPZj42iapodEK1FfodpSqaKZmynfxyd5nibln44HgUb0Tlc3PZT8qEcLy5VIG2wRoviaxQ/132\"}"
	bytes, code, err := users.CreateUser(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func TestGetUinfos() {
	var users = new(user.UserServer)
	var info = make(map[string]interface{})
	info["accids"] = "[\"101304574\"]"
	bytes, code, err := users.GetUinfos(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestCreateChatromm() {
	var chatroom = new(chatroom.ChatroomServer)
	var info = make(map[string]interface{})
	info["creator"] = "1234aaab"
	info["name"] = "拼字"
	info["announcement"] = "在线拼字"
	bytes, code, err := chatroom.CreateChatroom(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

var users = new(user.UserServer)

func GetUinfos() {
	var info = make(map[string]interface{})
	info["accids"] = "[100038367]"
	bytes, code, err := users.GetUinfos(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func UpdateUserUinfo() {
	var info = make(map[string]interface{})
	info["accid"] = "100038367"
	info["icon"] = "http://img.chenggua.com/Fl0Oai-JOC_xrcCbbFjKftQcL9K-"
	bytes, code, err := users.UpdateUserUinfo(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func UpdateTeamGroupInfo() {

	var info = make(map[string]interface{})
	info["tid"] = "2569132195"
	info["owner"] = "101405885"

	// info["joinmode"] = 1
	info["teamMemberLimit"] = 50

	logs.Debug("打印更新群请求体[%v]", info)
	bytes, _, err := teams.UpdateTeamInfo(info)
	logs.Error("-------", string(bytes))
	if err != nil {
		logs.Error("更新群信息异常,错误信息[%v]", err)
		return
	}
	return
}
