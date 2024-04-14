package ctype

import "encoding/json"

type ImagesStorgeType int

const (
	Local ImagesStorgeType = 1
	QiNiu ImagesStorgeType = 2
)

func (s ImagesStorgeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (r ImagesStorgeType) String() string {
	var str string
	switch r {
	case Local:
		str = "local"
	case QiNiu:
		str = "qiniu"
	default:
		str = "其他"
	}
	return str
}
