/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 11:30:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 22:44:41
 */

package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TechCatsLab/storage/mysql"
	"github.com/morgances/matchmaking/backend/conf"
	"strings"
)

var (
	DB *sql.DB
)

func init() {
	dsn := conf.MatchMakeConf.UserName + ":" + conf.MatchMakeConf.Password + "@" + conf.MatchMakeConf.Protocol +
		"(" + conf.MatchMakeConf.MysqlAddress + ":" + conf.MatchMakeConf.MysqlPort + ")/?parseTime=true"
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
			phone VARCHAR(20) DEFAULT 'NULL',
			wechat VARCHAR(20) DEFAULT 'NULL',
			nick_name VARCHAR(10),
			avatar VARCHAR(100),
			real_name VARCHAR(10) DEFAULT 'NULL',
			sex	TINYINT(1),
			birthday DATE DEFAULT '2018-1-1',
			height VARCHAR(10) DEFAULT 'NULL',
			location VARCHAR(10) DEFAULT 'NULL',
			job VARCHAR(20) DEFAULT 'NULL',
			faith VARCHAR(20) DEFAULT 'NULL',
			constellation VARCHAR(10) DEFAULT 'NULL',
			self_introduction VARCHAR(255) DEFAULT 'NULL',
			selec_criteria VARCHAR(255) DEFAULT 'NULL',

			open_id VARCHAR(255),
			age TINYINT UNSIGNED DEFAULT 0,
			create_at DATE,
			certified TINYINT(1) DEFAULT 0,
			vip TINYINT(1) DEFAULT 0,
			date_privilege INT UNSIGNED DEFAULT 0,
			points INT UNSIGNED DEFAULT 0,
			rose INT UNSIGNED DEFAULT 0,
			charm INT UNSIGNED DEFAULT 0,
			PRIMARY KEY(open_id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS admin(
			account VARCHAR(15),
			password VARCHAR(255),
			PRIMARY KEY (account))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`INSERT INTO admin(account,password) VALUES(?,?)`,
		conf.MatchMakeConf.AdmAccount, conf.MatchMakeConf.AdmPassword,
	)
	if err != nil {
		if !strings.Contains(err.Error(), "Duplicate entry") {
			return err
		}
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS post(
			id BIGINT AUTO_INCREMENT,
			open_id VARCHAR(255),
			title VARCHAR(50),
			content VARCHAR(255) DEFAULT 'NULL',
			date_time DATETIME,
			commend INT UNSIGNED DEFAULT 0,
			reviewed TINYINT DEFAULT 0,
			PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS follow(
			fan VARCHAR(255),
			idol VARCHAR(255),
			PRIMARY KEY (fan,idol))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS Signinin_record(
			open_id VARCHAR(255),
			signin_date DATE,
			PRIMARY KEY (open_id, signin_date))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS goods(
			id INT AUTO_INCREMENT,
			title VARCHAR(50) UNIQUE,
			price INT UNSIGNED,
			description VARCHAR(255) DEFAULT 'NULL',
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
			goods_id INT UNSIGNED,
			buyer_name VARCHAR(10),
			goods_title VARCHAR(50),
			date_time DATETIME,
			cost INT UNSIGNED,
			finished TINYINT DEFAULT 0,
			PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8
		`,
	)
	return err
}
