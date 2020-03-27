package parseconf

var Cfg *Config

type Global struct {
	Host string `conf:"host"`
	Port int64  `conf:"port"`
}

type LimitConfig struct {
	Count  int `conf:"count"`
	Second int `conf:"second"`
}

type RedisConfig struct {
	Addr     string `conf:"addr"`
	RPort    int64  `conf:"port"`
	Password string `conf:"password"`
	DB       int    `conf:"db"`
}

type MysqlConfig struct {
	MHost    string `conf:"host"`
	MPort    int64  `conf:"port"`
	UserName string `conf:"username"`
	PassWord string `conf:"password"`
	DBName   string `conf:"dbname"`
}

// Config 配置文件结构体
type Config struct {
	Global      `conf:"global"`
	LimitConfig `conf:"limit"`
	RedisConfig `conf:"redis"`
	MysqlConfig `conf:"mysql"`
}

func checkConf() {
	//这里做一些判断
}

func InitConf(confName string) (err error) {
	Cfg = &Config{}
	err = parseConf(confName, Cfg)
	checkConf()
	return
}
