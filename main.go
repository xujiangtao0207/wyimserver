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
	// UpdateTeamGroupInfo("2559483930", "101269535")
	TestGetTeamDetail("2610776150")
	// TestKickUserToTeam()
	// KickAllUserToTeam()
	// UpdateTeamAnnouncement(101404418, "", "2611391240")
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
	info["accid"] = fmt.Sprintf("%d", 101378960)
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
	info["owner"] = "100000001"
	//101374698
	info["tid"] = "2605758235"
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

func UpdateTeamGroupInfo(tid, owner string) {

	var info = make(map[string]interface{})
	info["tid"] = tid
	info["owner"] = owner

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

type RoomDetailRequest struct {
	RoomUuid      string `json:"room_uuid"`       //群tid
	RoomNum       string `json:"room_num"`        //房间码
	Parameter     int    `json:"parameter"`       //字数/时间
	GlobalStatus  int    `json:"global_status"`   //公有/私有 0拒绝 1 允许
	PeopleCount   int    `json:"people_count"`    //人数限制
	RoomDescribe  string `json:"room_describe"`   //房间描述
	RoomName      string `json:"room_name"`       //房间名称
	RoomType      int    `json:"room_type"`       //房间类型
	Status        int    `json:"status"`          //房间状态
	IsExistSecret int    `json:"is_exist_secret"` //是否存在密码
	Secret        string `json:"secret"`          //密码
}

var CommonInfoMap = map[string]string{
	"2560393181": "满腹经纶",
	"2560387530": "九天揽月",
	"2560377941": "日进斗金",
	"2560381653": "乘风破浪",
	"2560383622": "才高八斗",
	"2578282280": "学富五车",
}

func UpdateCommonRoomInfo() {
	// var uid = 100000001

	for key, value := range CommonInfoMap {
		//查询群信息
		_, _, custom, _, err := TestGetTeamDetail(key)
		var roomReq RoomDetailRequest
		err = json.Unmarshal([]byte(custom), &roomReq)
		if err != nil {
			fmt.Println("解析异常")
			return
		}
		fmt.Println(value)
		// roomReq.RoomNum = value
		// bytes, _ := json.Marshal(&roomReq)
		// err = UpdateTmpTeamGroupInfo(uid, 0, key, string(bytes))
		// if err != nil {
		// 	//创建群异常,销毁群操作
		// 	return
		// }
	}

}
func UpdateTmpTeamGroupInfo(uid, room_status int, roomUuid string, custom string) (err error) {

	var info = make(map[string]interface{})
	info["tid"] = roomUuid
	info["owner"] = fmt.Sprintf("%d", uid)

	var customInter map[string]interface{}
	err = json.Unmarshal([]byte(custom), &customInter)
	if err != nil {
		logs.Error("解析数据异常,错误信息[%v]", err)

	}

	if _, ok := customInter["room_status"]; ok {
		customInter["room_status"] = room_status
	} else {
		customInter["room_status"] = room_status
	}

	//不允许任何用户加入
	if room_status == 1 {
		info["joinmode"] = 2
	}

	bytesss, _ := json.Marshal(&customInter)

	info["custom"] = string(bytesss)

	logs.Debug("打印更新群请求体[%v]", info)
	bytes, _, err := teams.UpdateTeamInfo(info)

	if err != nil {
		logs.Error("更新群信息异常,错误信息[%v]", err)
		return err
	}

	logs.Debug("更新群信息响应体[%s],", string(bytes))
	return nil
}

func UpdateTeamAnnouncement(uid int, announcement, roomUuid string) {

	var info = make(map[string]interface{})
	info["tid"] = roomUuid
	info["owner"] = fmt.Sprintf("%d", uid)

	info["announcement"] = announcement
	logs.Debug("打印更新群请求体[%v]", info)
	bytes, _, err := teams.UpdateTeamInfo(info)

	if err != nil {
		logs.Error("更新群公告信息异常,错误信息[%v]", err)
		return
	}
	logs.Debug("更新群信息响应体[%s],", string(bytes))
}
