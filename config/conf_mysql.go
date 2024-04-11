package config

import "strconv"

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Config    string `yaml:"config"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_Level string `yaml:"log_level"` //日志等级，debug就是全部输出sql
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")" + "/" + m.Db + "?" + m.Config
	//root:1310138359@tcp(127.0.0.1:3306)/gvb_db?charset=utf8mb4&parseTime=True
}
