package configs

import (
	"github.com/spf13/viper"
)

var config = new(Config)

// Config 结构体用于存储配置信息
type Config struct {
	Bot                 BotConfig
	MySQL               MysqlConfig
	Redis               RedisConfig
	CustomPrefixPrompts []string
	CheckinAward        CheckinAwardConfig
}

// CheckinAwardConfig 结构体用于Bot配置信息
type CheckinAwardConfig struct {
	Max int64
	Min int64
}

// BotConfig 结构体用于Bot配置信息
type BotConfig struct {
	Token      string
	PublicURL  string
	ListenPort int
	Secret     string
}

// MySQLConfig 结构体用于存储数据库配置信息
type MysqlConfig struct {
	Addr     string
	Name     string
	Username string
	Password string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// loadConfig 函数用于加载配置文件
func init() {
	//导入配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile("configs/configs.yaml")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
	// fmt.Println(config)
}

func Get() Config {
	return *config
}
