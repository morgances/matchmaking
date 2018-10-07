package conf

import (
	"github.com/spf13/viper"
)

type matchMakeConfig struct {
	UserName string
	Password string
	Protocol string
	Address  string
	Port     string
	Database string
}

var MatchMakeConf *matchMakeConfig

func init() {
	viper.AddConfigPath("../conf")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	MatchMakeConf = &matchMakeConfig{
		UserName: viper.GetString("user_name"),
		Password: viper.GetString("password"),
		Protocol: viper.GetString("protocol"),
		Address:  viper.GetString("host.address"),
		Port:     viper.GetString("host.port"),
		Database: viper.GetString("database"),
	}
}
