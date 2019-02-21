package main

import (
	"github.com/astaxie/beego/logs"
	"wyimserver/chatroom"
	"wyimserver/team"
	"wyimserver/user"
)

func main() {
	TestCreateUser()
	TestAddUserToTeam()
	TestChangeTeamOwner()
}

func TestCreateTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["owner"] = "1234aaab"
	info["tname"] = "xujiangtao-team"
	info["members"] = "[\"1234aaaa\",\"1234aaab\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 0
	info["joinmode"] = 0
	bytes, code, err := teams.CreateTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

func TestAddUserToTeam() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = 1602192891
	info["owner"] = "1234aaabaaa"
	info["members"] = "[\"1234aaab\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 0
	info["joinmode"] = 0
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
	info["tid"] = 1602192891
	info["owner"] = "1234aaab"
	info["members"] = "[\"1234aaabaaa\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 0
	info["joinmode"] = 0
	bytes, code, err := teams.KickUserToTeam(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}
func TestGetTeamDetail() {
	var teams = new(team.TeamServer)
	var info = make(map[string]interface{})
	info["tid"] = 1602192891
	info["owner"] = "1234aaab"
	info["members"] = "[\"1234aaabaaa\"]"
	info["msg"] = "欢迎你加入橙瓜拼字聊天群"
	info["magree"] = 0
	info["joinmode"] = 0
	bytes, code, err := teams.GetTeamDetail(info)
	logs.Info("请求响应消息[%s]", string(bytes))
	logs.Info("请求响应消息[%d]", code)
	if err != nil {
		logs.Error("创建用户异常,错误信息:%v", err)
	}
}

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

func TestCreateUser() {
	var users = new(user.UserServer)
	var info = make(map[string]interface{})
	info["accid"] = "1234aaabaaa"
	info["name"] = "xujiangtao4"
	bytes, code, err := users.CreateUser(info)
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
