/*
 * @Author: zhanghao
 * @Date: 2018-10-08 22:01:22
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:30:38
 */

package model

import (
	"database/sql"
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
			return ErrDuplicateEntry
		}
		return err
	}
	return nil
}

func (signInServPrvd) FindByOpenID(oid string) (sis []SignInRecord, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM sign_in_record WHERE open_id=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		sis = append(sis, SignInRecord{})
		err = rows.Scan(
			&sis[i].OpenID, &sis[i].Date,
		)
		if err != nil {
			return nil, err
		}
	}
	return sis, nil
}

func (signInServPrvd) DeleteAllRecord(oid string) error {
	_, err := DB.Exec(
		`DELETE FROM sign_in_record WHERE oid=?`,
		oid,
	)
	return err
}
