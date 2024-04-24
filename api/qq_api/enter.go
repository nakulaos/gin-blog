package qq_api

import "gvb_server/models/ctype"

type QQLoginApi struct {
}
type AuthInfo struct {
	NickName    string           `json:"nick_name"`
	Gender      string           `json:"gender"`
	FigureurlQQ string           `json:"figureurl_qq"`
	Token       string           `json:"token"`
	SignStatus  ctype.SignStatus `json:"sign_status"` //注册来源
}

func NewAuthLogin(loginType, code string) (authInfo AuthInfo, err error) {
	switch loginType {
	case "qq":
		qqinfo, err := NewQQLogin(code)
		if err != nil {
			return authInfo, err
		}
		authInfo.NickName = qqinfo.Nickname
		authInfo.Gender = qqinfo.Gender
		authInfo.FigureurlQQ = qqinfo.FigureurlQQ
		authInfo.Token = qqinfo.OpenID
		authInfo.SignStatus = ctype.SignQQ
	}
	return
}
