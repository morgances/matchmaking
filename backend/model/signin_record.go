/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 22:01:22
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:30:38
 */

package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type (
	signInServPrvd struct{}
)

var (
	SigninService signInServPrvd
)

func (signInServPrvd) Insert(oid string) error {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	_, err := DB.Exec(
		`INSERT INTO signin_record(open_id, signin_date)
					VALUES(?,?)`,
		oid, date,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil
		}
		return err
	}
	return nil
}

func (signInServPrvd) FindByOpenID(oid string) (dates []string, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT signin_date FROM signin_record WHERE open_id=? LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var signDate string
		err = rows.Scan(&signDate)
		if err != nil {
			return nil, err
		}
		dates = append(dates, signDate)
	}
	return dates, nil
}
