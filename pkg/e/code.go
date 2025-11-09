package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
)

var MsgFlags = map[int]string{
	Success:       "success",
	Error:         "error",
	InvalidParams: "invalid params",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
