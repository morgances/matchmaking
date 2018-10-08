/*
 * @Author: zhanghao
 * @Date: 2018-10-08 11:30:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-08 19:54:09
 */

package model

import (
	"database/sql"

	"github.com/TechCatsLab/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/morgances/matchmaking/backend/conf"
)

var (
	DB *sql.DB
)

func init() {
	dsn := conf.MatchMakeConf.UserName + ":" + conf.MatchMakeConf.Password + "@" + conf.MatchMakeConf.Protocol +
		"(" + conf.MatchMakeConf.Address + ":" + conf.MatchMakeConf.Port + ")/?parseTime=true"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = initDatabase(); err != nil {
		panic(err)
	}
	if err = initTable(); err != nil {
		panic(err)
	}
}

func initDatabase() error {
	err := mysql.CreateDatabaseIfNotExist(DB, conf.MatchMakeConf.Database)
	if err != nil {
		return err
	}
	if _, err = DB.Exec("USE " + conf.MatchMakeConf.Database); err != nil {
		return err
	}
	return nil
}

func initTable() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS user(
		phone VARCHAR(20),
		wechat VARCHAR(20) UNIQUE,
		nick_name VARCHAR(10),
		avatar VARCHAR(100),
		real_name VARCHAR(10),
		sex	TINYINT(1),
		birthday DATE,
		height VARCHAR(10),
		location VARCHAR(10),
		job VARCHAR(20),
		faith VARCHAR(20),
		constellation VARCHAR(10),
		self_introduction VARCHAR(255),
		selec_criteria VARCHAR(255),

		open_id VARCHAR(255),
		create_at DATE,
		password VARCHAR(255),
		album VARCHAR(255),
		certified TINYINT(1),
		vip TINYINT(1),
		date_privilege INT,
		points INT,
		rose INT,
		charm INT,
		PRIMARY KEY(open_id)
	);`)
	return err
}
