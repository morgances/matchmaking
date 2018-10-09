/*
 * @Author: zhanghao
 * @Date: 2018-10-08 21:35:53
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

	Follow struct {
		Following string
		Followed  string
	}
)

var (
	FollowService followServPrvd
)

func (followServPrvd) Insert(f *Follow) error {
	exist, err := UserService.UserExist(f.Following)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotExist
	}
	exist, err = UserService.UserExist(f.Followed)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotExist
	}
	_, err = DB.Exec(
		`INSERT INTO follow(following, followed)
			VALUES(?,?)
		`,
		f.Following, f.Followed,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ErrDuplicateEntry
		}
		return err
	}
	return nil
}

func (followServPrvd) FindFollowing(oid string) (fs []Follow, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM follow WHERE following=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		fs = append(fs, Follow{})
		if err = rows.Scan(
			&fs[i].Following, &fs[i].Followed,
		); err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func (followServPrvd) FindFollowers(oid string) (fs []Follow, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM follow WHERE followed=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		fs = append(fs, Follow{})
		if err = rows.Scan(
			&fs[i].Following, &fs[i].Followed,
		); err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func (followServPrvd) Unfollow(following, followed string) error {
	result, err := DB.Exec(
		`DELETE FROM follow WHERE following=? AND followed=? LIMIT 1`,
		following, followed,
	)
	if err != nil {
		return err
	}
	var affected int64
	if affected, err = result.RowsAffected(); err != nil {
		return err
	}
	if affected == 0 {
		return ErrUnfollowFailed
	}
	return nil
}

func (followServPrvd) DeleteAllRecord(oid string) error {
	_, err := DB.Exec(
		`DELETE FROM follow WHERE following=?`,
		oid,
	)
	return err
}
