package errmsg

const (
	//status codes
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... 用户模块错误
	ERROR_UserName_Used    = 1001
	ERROR_Password_WRONG   = 1002
	ERROR_User_Not_Existed = 1003
	ERROR_Token_Not_Exist  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_FORMAT     = 1007
	//code = 2000... 文章模块错误

	//code = 3000... 分类模块错误
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "Failed",
	ERROR_UserName_Used:    "用户名已存在",
	ERROR_Password_WRONG:   "密码错误",
	ERROR_User_Not_Existed: "用户不存在",
	ERROR_Token_Not_Exist:  "Token不存在",
	ERROR_TOKEN_RUNTIME:    "Token过期",
	ERROR_TOKEN_WRONG:      "Token错误",
	ERROR_TOKEN_FORMAT:     "Token格式错误",
}

func Get_Error_Msg(code int) string {
	return codemsg[code]
}
