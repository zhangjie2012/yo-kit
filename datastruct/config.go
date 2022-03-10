package datastruct

type MySQLConf struct {
	Addr     string `yaml:"addr" json:"addr"`
	Database string `yaml:"database" json:"database"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type MongoConf struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type RedisConf struct {
	Addr     string `yaml:"addr" json:"addr"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

// WXMPConf 微信小程序配置
type WXMPConf struct {
	AppID  string `yaml:"appid"`
	Secret string `yaml:"secret"`
}
