/*
 * @Author: zhanghao
 * @Date: 2018-10-08 21:35:53
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-08 21:59:53
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
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
	_, err := DB.Exec(
		`INSERT INTO follow(following, followed)
			VALUES(?,?)
		`,
		f.Following, f.Followed,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New(fmt.Sprintf("duplicate entry -> %s follow %s", f.Following, f.Followed)) // need fix when struct field changed
		}
		return err
	}
	return nil
}

func (followServPrvd) FindFollowing(oid string) (fs []*Follow, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM follow WHERE following=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cols []string
	cols, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	fs = make([]*Follow, len(cols))
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(
			&fs[i].Following, &fs[i].Followed,
		)
		if err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func (followServPrvd) FindFollowers(oid string) (fs []*Follow, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM follow WHERE followed=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cols []string
	cols, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	fs = make([]*Follow, len(cols))
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(
			&fs[i].Following, &fs[i].Followed,
		)
		if err != nil {
			return nil, err
		}
	}
	return fs, nil
}

func (followServPrvd) Unfollow(oid string) error {
	_, err := DB.Exec(
		`DELETE FROM follow WHERE oid=? LIMIT 1`,
		oid,
	)
	return err
}
