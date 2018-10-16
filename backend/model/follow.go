/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 21:35:53
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:36:07
 */

package model

import (
	"database/sql"
	"strings"
)

type (
	followServPrvd struct{}
)

var (
	FollowService followServPrvd
)

func (followServPrvd) Insert(fan, idol string) error {
	_, err := DB.Exec(
		`INSERT INTO follow(fan, idol)
					VALUES(?,?)`,
		fan, idol,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ErrDuplicateEntry
		}
		return err
	}
	return nil
}

func (followServPrvd) FindFollowing(oid string) (us []User, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT u.open_id,u.nick_name
			  		FROM follow f JOIN user u ON f.idol=u.open_id 
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
			  		FROM follow f JOIN user u ON f.fan=u.open_id 
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
		`DELETE FROM follow WHERE fan=? AND idol=? LIMIT 1`,
		fan, idol,
	)
	return err
}

func (followServPrvd) FollowExist(fan, idol string) (bool, error) {
	row := DB.QueryRow(
		`SELECT COUNT(0) FROM follow WHERE fan=? AND idol=? LOCK IN SHARE MODE`,
		fan, idol,
	)
	var exist int32
	if err := row.Scan(&exist); err != nil {
		return false, err
	}
	return exist != 0, nil
}
