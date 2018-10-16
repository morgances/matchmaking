package conf

import (
	"github.com/spf13/viper"
)

type matchMakeConfig struct {
	// mysql config
	UserName     string
	Password     string
	Protocol     string
	MysqlAddress string
	MysqlPort    string
	Database     string

	// server config
	Address string
	Port    string

	PrivateTokenKey string
	AppID           string
	AppSecret       string

	// admin config
	AdmAccount  string
	AdmPassword string

	PreAlbumDir string
	AlbumSha1   string
}

var MatchMakeConf *matchMakeConfig

func init() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	MatchMakeConf = &matchMakeConfig{
		UserName:        viper.GetString("mysql.user_name"),
		Password:        viper.GetString("mysql.password"),
		Protocol:        viper.GetString("mysql.protocol"),
		MysqlAddress:    viper.GetString("mysql.host.address"),
		MysqlPort:       viper.GetString("mysql.host.port"),
		Database:        viper.GetString("mysql.database"),
		Address:         viper.GetString("host.address"),
		Port:            viper.GetString("host.port"),
		PrivateTokenKey: viper.GetString("private_token_key"),
		AppID:           viper.GetString("app_id"),
		AppSecret:       viper.GetString("app_secret"),
		AdmAccount:      viper.GetString("admin.account"),
		AdmPassword:     viper.GetString("admin.password"),
		PreAlbumDir:     viper.GetString("pre_album_dir"),
		AlbumSha1:       viper.GetString("album_sha1_secret"),
	}
}
