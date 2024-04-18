package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGitee SignStatus = 2
	SignEmail SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (r SignStatus) String() string {
	var str string
	switch r {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}
