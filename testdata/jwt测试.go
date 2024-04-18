package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/utils/jwt"
)

func main() {
	core.InitConf()
	core.InitLogger()

	token, err := jwt.GenToken(jwt.JwtPayLoad{
		Username: "zhangsan",
		Nickname: "dada",
		Role:     2,
		UserID:   1,
	})
	fmt.Println(token, err)

	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InpoYW5nc2FuIiwibmlja25hbWUiOiJkYWRhIiwicm9sZSI6MiwidXNlcl9pZCI6MSwiZXhwIjoxNzEzNDQ3MDQ1LCJpc3MiOiJ4eHgifQ.ZzvUsrapVipIKUqqkIucV9_tApl40S4VPABzvCuSkdo")
	fmt.Println(*claims, err)
}
