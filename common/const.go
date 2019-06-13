package common

const (
	// APPKEY    = "70a850cf0bd2006dea77401e9f489d1d"
	// APPSECRET = "18bc62ebe58f"

	APPKEY             = "60ed88b054f7476996388361ed9ad6f0"
	APPSECRET          = "5eda9eecc8c1"
	WyImEndpoint       = "https://api.netease.im/nimserver"
	CreateUserUrl      = "/user/create.action"      //创建用户
	UpdateUserUrl      = "/user/update.action"      // 网易云通信ID更新
	UpdateUserUinfoUrl = "/user/updateUinfo.action" //更新用户信息
	GetUinfosUrl       = "/user/getUinfos.action"   //获取用户名片

	CreateChatroomUrl = "/chatroom/create.action" //创建聊天室
	UpdateChatrommUrl = ""

	//群组
	CreateTeamUrl          = "/team/create.action"         //创建群组
	AddUserToTeamUrl       = "/team/add.action"            //拉人入群
	UpdateTeamInfoUrl      = "/team/update.action"         //编辑群资料
	KickUserToTeamUrl      = "/team/kick.action"           //踢人出群
	RemoveTeamUrl          = "/team/remove.action"         //解散群
	GetJoinTeamsForUserUrl = "/team/joinTeams.action"      //获取某用户所加入的群信息
	GetTeamDetailUrl       = "/team/queryDetail.action"    //查询指定群的详细信息
	LeaveTeamUrl           = "/team/leave.action"          //主动退群
	MuteTlistUrl           = "/team/muteTlist.action"      //禁言群成员
	MuteTlistAllUrl        = "/team/muteTlistAll.action"   //将群组整体禁言
	GetlistTeamMuteUrl     = "/team/listTeamMute.action"   //获取群组禁言列表
	UpdateTeamNickUrl      = "/team/updateTeamNick.action" //修改群昵称
	ChangeTeamOwnerUrl     = "/team/changeOwner.action"    //移交群主

	//消息
	SendMsgUrl            = "/msg/sendMsg.action"            //发送自定义系统通知
	SendAttachMsgUrl      = "/msg/sendAttachMsg.action"      //发送自定义系统通知
	SendBatchAttachMsgUrl = "/msg/sendBatchAttachMsg.action" //批量发送点对点自定义系统通知
)
