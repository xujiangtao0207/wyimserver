package team

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"wyimserver/common"
	"wyimserver/utils"
)

type TeamServer struct {
	group TeamModule
}

type TeamModule interface {
	CreateTeam(map[string]interface{})          //创建群组
	AddUserToTeam(map[string]interface{})       //拉人入群
	UpdateTeamInfo(map[string]interface{})      //更新群信息
	KickUserToTeam(map[string]interface{})      //踢人出群
	RemoveTeam(map[string]interface{})          //解散群
	GetJoinTeamsForUser(map[string]interface{}) //获取某用户所加入的群信息
	GetTeamDetail(map[string]interface{})       //查询指定群的详细信息
	LeaveTeam(map[string]interface{})           //主动退群
	MuteTlist(map[string]interface{})           //禁言群成员
	MuteTlistAll(map[string]interface{})        //将群组整体禁言
	GetlistTeamMute(map[string]interface{})     //获取群组禁言列表
	UpdateTeamNick(map[string]interface{})      //修改群昵称
	ChangeTeamOwner(map[string]interface{})     //移交群主
}

func (group *TeamServer) CreateTeam(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.CreateTeamUrl, "POST", []byte(GetBodyStr(info)))
}
func (group *TeamServer) AddUserToTeam(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.AddUserToTeamUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) UpdateTeamInfo(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.UpdateTeamInfoUrl, "POST", []byte(GetBodyStr(info)))
}
func (group *TeamServer) KickUserToTeam(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.KickUserToTeamUrl, "POST", []byte(GetBodyStr(info)))
}
func (group *TeamServer) RemoveTeam(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.RemoveTeamUrl, "POST", []byte(GetBodyStr(info)))
}
func (group *TeamServer) GetJoinTeamsForUser(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.GetJoinTeamsForUserUrl, "POST", []byte(GetBodyStr(info)))
}
func (group *TeamServer) GetTeamDetail(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.GetTeamDetailUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) LeaveTeam(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.LeaveTeamUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) MuteTlist(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.MuteTlistUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) MuteTlistAll(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.MuteTlistAllUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) GetlistTeamMute(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.GetlistTeamMuteUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) UpdateTeamNick(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.UpdateTeamNickUrl, "POST", []byte(GetBodyStr(info)))
}

func (group *TeamServer) ChangeTeamOwner(info map[string]interface{}) ([]byte, int, error) {
	return utils.SendRequestToWy(common.WyImEndpoint+common.ChangeTeamOwnerUrl, "POST", []byte(GetBodyStr(info)))
}

func GetBodyStr(info map[string]interface{}) (bodyStr string) {
	if tname, ok := info["tname"]; ok {
		bodyStr += fmt.Sprintf("tname=%s", tname.(string))
	}

	if tid, ok := info["tid"]; ok {
		bodyStr = common.GetBodyPrefix(bodyStr)
		bodyStr += fmt.Sprintf("tid=%d", tid.(int))
	}

	if owner, ok := info["owner"]; ok {
		bodyStr += fmt.Sprintf("&owner=%s", owner.(string))
	}

	if nick, ok := info["nick"]; ok {
		bodyStr += fmt.Sprintf("&nick=%s", nick.(string))
	}

	if members, ok := info["members"]; ok {
		// var memberArr []string
		// for _, member := range members.([]interface{}) {
		// 	memberArr = append(memberArr, member.(string))
		// }
		bodyStr += fmt.Sprintf("&members=%s", members)
	}
	if announcement, ok := info["announcement"]; ok {
		bodyStr += fmt.Sprintf("&announcement=%s", announcement.(string))
	}
	if intro, ok := info["intro"]; ok {
		bodyStr += fmt.Sprintf("&intro=%s", intro.(string))
	}
	if msg, ok := info["msg"]; ok {
		bodyStr += fmt.Sprintf("&msg=%s", msg.(string))
	}
	if magree, ok := info["magree"]; ok {
		bodyStr += fmt.Sprintf("&magree=%d", magree.(int))
	}
	if joinmode, ok := info["joinmode"]; ok {
		bodyStr += fmt.Sprintf("&joinmode=%d", joinmode.(int))
	}

	if icon, ok := info["icon"]; ok {
		bodyStr += fmt.Sprintf("&icon=%s", icon.(string))
	}

	if custom, ok := info["custom"]; ok {
		bodyStr += fmt.Sprintf("&custom=%s", custom.(string))
	}

	if beinvitemode, ok := info["beinvitemode"]; ok {
		bodyStr += fmt.Sprintf("&beinvitemode=%d", beinvitemode.(int))
	}

	if invitemode, ok := info["invitemode"]; ok {
		bodyStr += fmt.Sprintf("&invitemode=%d", invitemode.(int))
	}

	if uptinfomode, ok := info["uptinfomode"]; ok {
		bodyStr += fmt.Sprintf("&uptinfomode=%d", uptinfomode.(int))
	}

	if upcustommode, ok := info["upcustommode"]; ok {
		bodyStr += fmt.Sprintf("&upcustommode=%d", upcustommode.(int))
	}

	if teamMemberLimit, ok := info["teamMemberLimit"]; ok {
		bodyStr += fmt.Sprintf("&teamMemberLimit=%d", teamMemberLimit.(int))
	}
	if accid, ok := info["accid"]; ok {
		bodyStr += fmt.Sprintf("&accid=%s", accid.(string))
	}
	//新群组帐号
	if newowner, ok := info["newowner"]; ok {
		bodyStr += fmt.Sprintf("&newowner=%s", newowner.(string))
	}
	//1:群主解除群主后离开群，2：群主解除群主后成为普通成员。其它414
	if leave, ok := info["leave"]; ok {
		bodyStr += fmt.Sprintf("&leave=%d", leave.(int))
	}
	//拉人入群组
	if attach, ok := info["attach"]; ok {
		bodyStr += fmt.Sprintf("&attach=%s", attach.(string))
	}
	logs.Info("打印body[%s]", bodyStr)
	return bodyStr
}
