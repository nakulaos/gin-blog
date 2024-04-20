package config

type Redis struct {
	Ip       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}
