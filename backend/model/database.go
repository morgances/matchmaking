/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 11:30:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 22:44:41
 */

package model

import (
	"database/sql"
	"strings"

	"github.com/morgances/matchmaking/backend/conf"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func init() {
	dsn := conf.MMConf.UserName + ":" + conf.MMConf.Password + "@" + conf.MMConf.Protocol +
		"(" + conf.MMConf.MysqlAddress + ":" + conf.MMConf.MysqlPort + ")/?parseTime=true"
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
	_, err := DB.Exec("CREATE DATABASE IF NOT EXISTS " + conf.MMConf.Database)
	if err != nil {
		return err
	}
	if _, err = DB.Exec("USE " + conf.MMConf.Database); err != nil {
		return err
	}

	return nil
}

func initTable() error {
	_, err := DB.Exec(
		`CREATE TABLE IF NOT EXISTS user(
					phone VARCHAR(20) DEFAULT 'null',
					wechat VARCHAR(20) DEFAULT 'null',
					nick_name VARCHAR(10),
					real_name VARCHAR(10) DEFAULT 'null',
					sex	TINYINT(1) UNSIGNED NOT NULL,
					birthday DATE DEFAULT '2018-01-01',
					height VARCHAR(10) DEFAULT 'null',
					location VARCHAR(10) DEFAULT 'null',
					job VARCHAR(20) DEFAULT 'null',
					faith VARCHAR(20) DEFAULT 'null',
					constellation VARCHAR(10) DEFAULT 'null',
					self_introduction VARCHAR(255) DEFAULT 'null',
					selec_criteria VARCHAR(255) DEFAULT 'null',

					open_id VARCHAR(35),
					age TINYINT UNSIGNED DEFAULT 0,
					create_at DATE,
					certified TINYINT(1) DEFAULT 0,
					vip TINYINT(1) DEFAULT 0,
					date_privilege INT UNSIGNED DEFAULT 0,
					points INT UNSIGNED DEFAULT 0,
					rose INT UNSIGNED DEFAULT 0,
					charm INT UNSIGNED DEFAULT 0,
					PRIMARY KEY(open_id))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS admin(
					account VARCHAR(15),
					password VARCHAR(255),
					PRIMARY KEY (account))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`INSERT INTO admin(account,password) VALUES(?,?)`,
		conf.MMConf.AdmAccount, conf.MMConf.AdmPassword,
	)
	if err != nil {
		if !strings.Contains(err.Error(), "Duplicate entry") {
			return err
		}
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS post(
					id BIGINT AUTO_INCREMENT,
					open_id VARCHAR(35) NOT NULL,
					title VARCHAR(50) NOT NULL,
					content VARCHAR(255) DEFAULT 'null',
					date_time DATETIME NOT NULL,
					commend INT UNSIGNED DEFAULT 0,
					reviewed TINYINT UNSIGNED DEFAULT 0,
					PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS follow(
					fan VARCHAR(255),
					idol VARCHAR(255),
					PRIMARY KEY (fan,idol))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS signin_record(
					open_id VARCHAR(35),
					signin_date DATE,
					PRIMARY KEY (open_id, signin_date))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS goods(
					id INT AUTO_INCREMENT,
					title VARCHAR(50) UNIQUE,
					price FLOAT(12,2) UNSIGNED NOT NULL,
					description VARCHAR(255) DEFAULT 'null',
					PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS trade(
					id BIGINT AUTO_INCREMENT,
					open_id VARCHAR(35) NOT NULL,
					goods_id INT UNSIGNED NOT NULL,
					buyer_name VARCHAR(10) NOT NULL,
					goods_title VARCHAR(50) NOT NULL,
					date_time DATETIME NOT NULL,
					cost FLOAT(12,2) UNSIGNED NOT NULL,
					finished TINYINT UNSIGNED DEFAULT 0,
					PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS recharge(
						id INT AUTO_INCREMENT,
						open_id VARCHAR(35) NOT NULL,
						project VARCHAR(20) NOT NULL,
						recharge_num INT UNSIGNED NOT NULL,
						fee INT UNSIGNED NOT NULL,
						transaction_id VARCHAR(255) DEFAULT 'null',
						status TINYINT UNSIGNED NOT NULL DEFAULT 0,
						PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8`,
	)
	return err
}
