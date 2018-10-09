/*
 * @Author: zhanghao
 * @Date: 2018-10-08 11:30:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 14:35:06
 */

package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TechCatsLab/storage/mysql"
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
	_, err := DB.Exec(
		`CREATE TABLE IF NOT EXISTS user(
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
			PRIMARY KEY(open_id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS post(
			id BIGINT AUTO_INCREMENT,
			open_id VARCHAR(255),
			title VARCHAR(50),
			image VARCHAR(255),
			content VARCHAR(255),
			date_time DATETIME,
			like INT,
			PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS follow(
			following VARCHAR(255),
			followed VARCHAR(255),
			PRIMARY KEY (following,followed))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS Sign_in_record(
			open_id VARCHAR(255),
			date DATE,
			PRIMARY KEY (open_id, date))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS goods(
			id INT AUTO_INCREMENT,
			name VARCHAR(50),
			price INT,
			description VARCHAR(255),
			image VARCHAR(255),
			PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS trade(
			id BIGINT AUTO_INCREMENT,
			open_id VARCHAR(255),
			goods_id INT,
			data_time DATETIME,
			cost INT,
			PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	return err
}
