/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 21:35:53
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:36:07
 */

package model

import (
	"database/sql"
	"github.com/morgances/matchmaking/backend/conf"
)

type (
	followServPrvd struct{}
)

var (
	FollowService followServPrvd
)

func (followServPrvd) Insert(fan, idol string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE `+conf.MMConf.Database+`.user SET rose=rose-1 WHERE open_id=? LIMIT 1`, fan)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		`INSERT INTO `+conf.MMConf.Database+`.follow(fan, idol) VALUES(?,?)`,
		fan, idol,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (followServPrvd) FindFollowing(oid string) (us []User, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT u.open_id,u.nick_name
			  		FROM `+conf.MMConf.Database+`.follow f JOIN `+conf.MMConf.Database+`.user u ON f.idol=u.open_id 
			  		WHERE f.fan=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var user User
		if err = rows.Scan(&user.OpenID, &user.NickName); err != nil {
			return nil, err
		}
		us = append(us, user)
	}
	return us, nil
}

func (followServPrvd) FindFollower(oid string) (us []User, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT u.open_id,u.nick_name
			  		FROM `+conf.MMConf.Database+`.follow f JOIN `+conf.MMConf.Database+`.user u ON f.fan=u.open_id 
			  		WHERE f.idol=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var user User
		if err = rows.Scan(&user.OpenID, &user.NickName); err != nil {
			return nil, err
		}
		us = append(us, user)
	}
	return us, nil
}

func (followServPrvd) Delete(fan, idol string) error {
	_, err := DB.Exec(
		`DELETE FROM `+conf.MMConf.Database+`.follow WHERE fan=? AND idol=? LIMIT 1`,
		fan, idol,
	)
	return err
}

func (followServPrvd) FollowExist(fan, idol string) (bool, error) {
	row := DB.QueryRow(
		`SELECT COUNT(0) FROM `+conf.MMConf.Database+`.follow WHERE fan=? AND idol=? LOCK IN SHARE MODE`,
		fan, idol,
	)
	var exist int32
	if err := row.Scan(&exist); err != nil {
		return false, err
	}
	return exist != 0, nil
}
