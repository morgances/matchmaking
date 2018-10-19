/*
 * Revision History:
 *     Initial: 2018/10/19        Zhang Hao
 */

package model

import (
	"database/sql"
	"errors"

	"github.com/morgances/matchmaking/backend/conf"
)

type (
	rechargeServPrvd struct{}

	Recharge struct {
		ID            int
		OpenID        string
		Project       string
		Num           int
		Fee           int
		TransactionID string
		Status        uint8
	}
)

var (
	RechargeService rechargeServPrvd
)

// Insert limit recharge project value to 'vip' and 'rose', num regard as 1 when proj='vip'
func (rechargeServPrvd) Insert(proj, openid string, num int) (id int, err error) {
	fee := 0
	switch proj {
	case "vip":
		fee = conf.MMConf.VIPFee
	case "rose":
		fee = num * 100
	default:
		return 0, errors.New("unknown recharge project: " + proj)
	}
	var rslt sql.Result
	rslt, err = DB.Exec(
		`INSERT INTO recharge(open_id,project,recharge_num,fee)
					VALUES(?,?,?,?)`,
		openid, proj, num, fee,
	)
	if err != nil || rslt == nil {
		return 0, errors.New("insert recharge record: " + err.Error())
	}
	var lastid int64
	lastid, err = rslt.LastInsertId()
	return int(lastid), err
}

func (rechargeServPrvd) FindAll() ([]Recharge, error) {
	rows, err := DB.Query(`SELECT * FROM recharge ORDER BY id DESC LOCK IN SHARE MODE`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rchgs []Recharge
	for i := 0; rows.Next(); i++ {
		rchg := Recharge{}
		err = rows.Scan(
			rchg.ID, rchg.OpenID, rchg.Project, rchg.Num, rchg.Fee, rchg.TransactionID, rchg.Status,
		)
		if err != nil {
			return nil, err
		}
		rchgs = append(rchgs, rchg)
	}
	return rchgs, nil
}

func (rechargeServPrvd) Success(id int, transid string) error {
	info, err := RechargeService.findByID(id)
	if err != nil {
		return errors.New("Success: " + err.Error())
	}
	switch info.Project {
	case "vip":
		return upgradeVip(info.ID, info.OpenID, transid)
	case "rose":
		return rechargeRose(info.ID, info.Num, info.OpenID, transid)
	default:
		return errors.New("Success: unknown recharge project " + info.Project)
	}
}

func (rechargeServPrvd) Fail(id int) error {
	_, err := DB.Exec(`UPDATE recharge SET status=2 WHERE id=? LIMIT 1`, id)
	return err
}

func (rechargeServPrvd) findByID(id int) (*Recharge, error) {
	row := DB.QueryRow(
		`SELECT * FROM recharge WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	rchg := &Recharge{}
	err := row.Scan(&rchg.ID, &rchg.OpenID, &rchg.Project, &rchg.Num, &rchg.Fee, &rchg.TransactionID, &rchg.Status)
	if err != nil {
		return nil, errors.New("findByID: " + err.Error())
	}
	return rchg, nil
}

func upgradeVip(id int, openid, transid string) error {
	var tx *sql.Tx
	tx, err := DB.Begin()
	if err != nil {
		return errors.New("upgradeVip: " + err.Error())
	}
	var rslt sql.Result
	rslt, err = tx.Exec(
		`UPDATE user SET vip=1,points=points+520,rose=rose+520,date_privilege=date_privilege+1 WHERE open_id=? LIMIT 1`,
		openid,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("upgradeVip: " + err.Error())
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("upgradeVip: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("upgradeVip: user may not exists")
	}
	rslt, err = tx.Exec(
		`UPDATE recharge SET status=1,transaction_id=? WHERE id=? AND status=0 LIMIT 1`,
		transid, id,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("upgradeVip: " + err.Error())
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("upgradeVip: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("upgradeVip: recharge may not exists, or is already handled")
	}
	return tx.Commit()
}

func rechargeRose(id, num int, openid, transid string) error {
	var tx *sql.Tx
	tx, err := DB.Begin()
	if err != nil {
		return errors.New("rechargeRose: " + err.Error())
	}
	var rslt sql.Result
	rslt, err = tx.Exec(
		`UPDATE user SET rose=rose+?,points=points+? WHERE open_id=? LIMIT 1`,
		num, num*10, openid,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("rechargeRose: " + err.Error())
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("rechargeRose: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("rechargeRose: user may not exists")
	}
	rslt, err = tx.Exec(
		`UPDATE recharge SET status=1,transaction_id=? WHERE id=? AND status=0 LIMIT 1`,
		transid, id,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("rechargeRose: " + err.Error())
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("rechargeRose: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("rechargeRose: recharge may not exists, or is already handled")
	}
	return tx.Commit()
}
