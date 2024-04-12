package res

type ErrorCode int

const (
	SettingsError  ErrorCode = 7771 //系统错误
	ParameterError ErrorCode = 7772 //参数错误
)

var (
	ErrMap = map[ErrorCode]string{
		SettingsError:  "系统错误",
		ParameterError: "参数错误",
	}
)
