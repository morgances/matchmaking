/*
 * @Author: zhanghao
 * @Date: 2018-10-08 22:01:22
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-08 22:17:45
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type (
	signInServPrvd struct{}

	SignInRecord struct {
		OpenID string
		Date   time.Time
	}
)

var (
	SignInService signInServPrvd
)

func (signInServPrvd) Insert(si *SignInRecord) error {
	_, err := DB.Exec(
		`INSERT INTO sign_in_record(open_id, date)
			VALUES(?,?)
		`,
		si.OpenID, si.Date,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New(fmt.Sprintf("duplicate entry -> %s signIn at %s", si.OpenID, si.Date)) // need fix when struct field changed
		}
		return err
	}
	return nil
}

func (signInServPrvd) FindByOpenID(oid string) (sis []*SignInRecord, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM sign_in_record WHERE open_id=? LOCK IN SHARE MODE`,
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
	sis = make([]*SignInRecord, len(cols))
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(
			&sis[i].OpenID, &sis[i].Date,
		)
		if err != nil {
			return nil, err
		}
	}
	return sis, nil
}
