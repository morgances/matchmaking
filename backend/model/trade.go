/*
 * @Author: zhanghao
 * @Date: 2018-10-09 10:36:35
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:39:18
 */

package model

import (
	"database/sql"
	"strings"
	"time"
)

type (
	tradeServPrvd struct{}

	Trade struct {
		ID       int64
		OpenID   string
		GoodsID  int64
		DateTime time.Time
		Cost     int64
	}
)

var (
	TradeService tradeServPrvd
)

func (tradeServPrvd) Insert(t *Trade) error {
	_, err := DB.Exec(
		`INSERT INTO trade(open_id,goods_id,data_time,cost)
			VALUES(?,?,?,?)
		`,
		t.OpenID, t.GoodsID, t.DateTime, t.Cost,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ErrDuplicateEntry
		}
		return err
	}
	return nil
}

func (tradeServPrvd) FindByID(id int64) (t *Trade, err error) {
	row := DB.QueryRow(
		`SELECT * FROM trade WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	if err = row.Scan(
		&t.ID, &t.OpenID, &t.GoodsID, &t.DateTime, &t.Cost,
	); err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return
}

func (tradeServPrvd) FindByOpenID() (ts []Trade, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM trade WHERE open_id=? ORDER BY date_time DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ts = append(ts, Trade{})
		err = rows.Scan(
			&ts[i].ID, &ts[i].OpenID, &ts[i].GoodsID, &ts[i].DateTime, &ts[i].Cost,
		)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}

func (tradeServPrvd) DeleteByID(id int64) error {
	_, err := DB.Exec(
		`DELETE FROM trade WHERE id=? LIMIT 1`,
		id,
	)
	return err
}
