package models

//统计用户登录数据 id,用户id，用户昵称，用户token，登陆设备，登陆时间

type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` //登陆的ip
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登陆设备
	Addr      string    `gorm:"size:128" json:"addr"`
}
