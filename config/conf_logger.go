package config

type Logger struct {
	Level          string `yaml:"level"`
	Prefix         string `yaml:"prefix"`
	Director       string `yaml:"director"`
	Show_Line      bool   `yaml:"show_line"`      //是否显示行号
	Log_In_Console bool   `yaml:"log_in_console"` //是否显示打印的路径
}
