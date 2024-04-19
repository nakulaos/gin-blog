package utils

import "strings"

// 13362475924
// 133****5924
func DesensitizationTel(tel string) string {
	if len(tel) != 11 {
		return "无效号码"
	}
	return tel[:3] + "****" + tel[7:]
}

// 20491002@163.com
// 2*******@163.com
func DesensitizationEmail(email string) string {

	emailList := strings.Split(email, "@")
	if len(emailList) != 2 {
		return "无效邮箱"
	}
	var musk string
	for i := 0; i < len(emailList[0])-1; i++ {
		musk += "*"
	}
	return emailList[0][:1] + musk + "@" + emailList[1]
}
