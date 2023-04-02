package errmsg

const (
	SUCCESS              = 200
	ERROR                = 500
	error_username_used  = 1001
	error_password_wrong = 1002
	//error_password_used=2001
	error_user_notexist  = 1003
	error_token_exist    = 1004 //TOKEN?
	error_token_timeout  = 1005 //
	error_token_wrong    = 1006 //?
	error_token_type     = 1007
	ERROR_ID_USED        = 2001
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codemsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	error_username_used:  "用户名已被使用",
	error_password_wrong: "密码错误",
	//error_password_used:"密码已被使用",
	error_user_notexist: "用户名不存在",
	error_token_exist:   "TOKEN 不存在", //TOKEN?
	error_token_timeout: "TOKEN 已过期", //
	error_token_wrong:   "TOKEN 错误",
	error_token_type:    "TOKEN 格式错误",
	ERROR_ID_USED:       "此ID已被使用",

	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
