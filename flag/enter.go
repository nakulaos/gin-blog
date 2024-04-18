package flag

import sys_flag "flag"

type Option struct {
	DB   bool
	User string
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop db为true，则停止web项目，只是数据迁移
func IsWebStop(option Option) bool {
	if option.DB || option.User != "" {
		return true
	}
	return false
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	is_show := true
	if option.DB {
		MakeMigrations()
		is_show = false
	}
	if option.User != "" {
		CreateUser(option.User)
		is_show = false
	}

	if is_show {
		sys_flag.Usage()
	}
}
