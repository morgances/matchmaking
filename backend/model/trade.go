/*
 * @Author: zhanghao
 * @DateTime: 2018-10-09 10:36:35
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:39:18
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/morgances/matchmaking/backend/conf"
)

var (
	ErrMakeTrade = errors.New("error make a trade")
)

type (
	tradeServPrvd struct{}

	Trade struct {
		ID        int64
		OpenID    string
		GoodsID   int64
		BuyerName string
		GoodsName string
		DateTime  time.Time
		Cost      float64
		Finished  bool
	}
)

var (
	TradeService tradeServPrvd
)

func (tradeServPrvd) Insert(t *Trade) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	var rslt sql.Result
	if rslt, err = tx.Exec(`UPDATE `+conf.MMConf.Database+`.user SET points=points-? WHERE open_id=? LIMIT 1`, t.Cost, t.OpenID); err != nil {
		tx.Rollback()
		return err
	}

	if affec, err := rslt.RowsAffected(); err == nil && affec != 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("maybe user(%s): %s not exist", t.BuyerName,t.OpenID))
	}

	_, err = DB.Exec(
		`INSERT INTO `+conf.MMConf.Database+`.trade(open_id,goods_id,buyer_name,goods_title,cost,date_time,finished)
					VALUES(?,?,?,?,?,NOW(),0)`,
		t.OpenID, t.GoodsID, t.BuyerName, t.GoodsName, t.Cost,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Cancel can only cancel unfinished order
func (tradeServPrvd) Cancel(t *Trade) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	var rslt sql.Result
	if rslt, err = tx.Exec(`UPDATE `+conf.MMConf.Database+`.user SET points=points+? WHERE open_id=? LIMIT 1`, t.Cost, t.OpenID); err != nil {
		tx.Rollback()
		return errors.New("failed to return points to user(id): "+t.OpenID)
	}
	if affec, err := rslt.RowsAffected(); err == nil && affec != 1 {
		tx.Rollback()
		return errors.New("failed to return points to user(id): "+t.OpenID)
	}
	rslt, err = DB.Exec(
		`DELETE FROM `+conf.MMConf.Database+`.trade WHERE id=? AND finished=0 LIMIT 1`,
		t.ID,
	)
	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("failed to delete trade record id: %d", t.ID))
	}
	if affec, err := rslt.RowsAffected(); err == nil && affec != 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("maybe trade record id:%d not exist", t.ID))
	}
	return tx.Commit()
}

func (tradeServPrvd) FindByID(id int64) (*Trade, error) {
	row := DB.QueryRow(
		`SELECT * FROM `+conf.MMConf.Database+`.trade WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	t := Trade{}
	if err := row.Scan(
		&t.ID, &t.OpenID, &t.GoodsID, &t.BuyerName, &t.GoodsName, &t.DateTime, &t.Cost, &t.Finished,
	); err != nil {
		return nil, err
	}
	return &t, nil
}

func (tradeServPrvd) FindByOpenID(oid string) (ts []Trade, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM `+conf.MMConf.Database+`.trade WHERE open_id=? ORDER BY date_time DESC LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ts = append(ts, Trade{})
		err = rows.Scan(
			&ts[i].ID, &ts[i].OpenID, &ts[i].GoodsID, &ts[i].BuyerName, &ts[i].GoodsName, &ts[i].DateTime, &ts[i].Cost, &ts[i].Finished,
		)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}

func (tradeServPrvd) FindUnfinishedTrade() (ts []Trade, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM ` + conf.MMConf.Database + `.trade WHERE finished=0 ORDER BY date_time DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ts = append(ts, Trade{})
		err = rows.Scan(
			&ts[i].ID, &ts[i].OpenID, &ts[i].GoodsID, &ts[i].BuyerName, &ts[i].GoodsName, &ts[i].DateTime, &ts[i].Cost, &ts[i].Finished,
		)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}

func (tradeServPrvd) UpdateTradeStatus(id int64) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.trade SET finished=1 WHERE id=? LIMIT 1`,
		id,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update status of trade: %d", id))
	}
	return nil
}
