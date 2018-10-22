/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
 */

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

	// app config
	AppID           string
	AppSecret       string
	MchID           string
	AppOrderKey     string
	VIPFee          uint32
	PrivateTokenKey string

	// admin config
	AdmAccount  string
	AdmPassword string
}

var MMConf *matchMakeConfig

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	MMConf = &matchMakeConfig{
		UserName:     viper.GetString("mysql.user_name"),
		Password:     viper.GetString("mysql.password"),
		Protocol:     viper.GetString("mysql.protocol"),
		MysqlAddress: viper.GetString("mysql.host.address"),
		MysqlPort:    viper.GetString("mysql.host.port"),
		Database:     viper.GetString("mysql.database"),

		Address: viper.GetString("host.address"),
		Port:    viper.GetString("host.port"),

		PrivateTokenKey: viper.GetString("app.private_token_key"),
		AppID:           viper.GetString("app.app_id"),
		AppSecret:       viper.GetString("app.app_secret"),
		MchID:           viper.GetString("app.mch_id"),
		VIPFee:          uint32(viper.GetInt32("app.vip_fee")),
		AppOrderKey:     viper.GetString("app.key"),

		AdmAccount:  viper.GetString("admin.account"),
		AdmPassword: viper.GetString("admin.password"),
	}
}
