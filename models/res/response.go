package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 666
	Error   = 777
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功!", c)

}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)

}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)

}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrMap[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(int(code), map[string]any{}, "未知错误，该错误码没有对应错误描述！", c)
}
