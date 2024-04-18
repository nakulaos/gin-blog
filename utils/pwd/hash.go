package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"gvb_server/global"
)

// 加密密码
func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Info(err)
	}
	return string(hash)
}

// 验证密码 hash密码
func ValidPassword(CurrentHashPassword, RawPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(CurrentHashPassword), []byte(RawPassword)); err != nil {
		global.Log.Info(err)
		return false
	}
	return true

}
