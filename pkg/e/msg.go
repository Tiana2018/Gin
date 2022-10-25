package e
var MsgFlags = map[int]string {
	SUCCESS : "ok", // 200
	ERROR : "fail", // 500
	INVALID_PARAMS : "请求参数错误", // 400
	ERROR_EXIST_TAG : "已存在该标签名称", // 10001
	ERROR_NOT_EXIST_TAG : "该标签不存在", // 10002
	ERROR_NOT_EXIST_ARTICLE : "该文章不存在", // 10003
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败", // 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时", // 20002
	ERROR_AUTH_TOKEN : "Token生成失败", // 20003
	ERROR_AUTH : "Token错误", // 20004
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
