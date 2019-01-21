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
		ID            uint32
		OpenID        string
		Project       string
		Num           uint32
		Fee           uint32
		TransactionID string
		Status        uint8
	}
)

var (
	RechargeService rechargeServPrvd
)

// Insert limit recharge project value to 'vip' and 'rose', num regard as 1 when proj='vip'
func (rechargeServPrvd) Insert(proj, openid string, num uint32) (id uint32, err error) {
	var fee uint32 = 0
	switch proj {
	case "vip":
		fee = conf.MMConf.VIPFee
	case "rose":
		fee = num * conf.MMConf.RoseFee
	default:
		return 0, errors.New("unknown recharge project: " + proj)
	}
	var rslt sql.Result
	rslt, err = DB.Exec(
		`INSERT INTO `+conf.MMConf.Database+`.recharge(open_id,project,recharge_num,fee)
					VALUES(?,?,?,?)`,
		openid, proj, num, fee,
	)
	if err != nil || rslt == nil {
		return 0, errors.New("insert recharge record: " + err.Error())
	}
	var lastid int64
	lastid, err = rslt.LastInsertId()
	return uint32(lastid), err
}

func (rechargeServPrvd) FindAll() ([]Recharge, error) {
	rows, err := DB.Query(`SELECT * FROM ` + conf.MMConf.Database + `.recharge ORDER BY id DESC LOCK IN SHARE MODE`)
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
	return rchgs, rows.Err()
}

func (rechargeServPrvd) Success(id uint32, transid string) error {
	info, err := RechargeService.findByID(id)
	if err != nil {
		return err
	}
	switch info.Project {
	case "vip":
		return upgradeVip(info.ID, info.OpenID, transid)
	case "rose":
		return rechargeRose(info.ID, info.Num, info.OpenID, transid)
	default:
		return errors.New("unknown recharge project: " + info.Project)
	}
}

func (rechargeServPrvd) Fail(id uint32) error {
	_, err := DB.Exec(`UPDATE `+conf.MMConf.Database+`.recharge SET status=2 WHERE id=? LIMIT 1`, id)
	return err
}

func (rechargeServPrvd) findByID(id uint32) (*Recharge, error) {
	row := DB.QueryRow(
		`SELECT * FROM `+conf.MMConf.Database+`.recharge WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	rchg := &Recharge{}
	err := row.Scan(&rchg.ID, &rchg.OpenID, &rchg.Project, &rchg.Num, &rchg.Fee, &rchg.TransactionID, &rchg.Status)
	if err != nil {
		return nil, errors.New("findByID: " + err.Error())
	}
	return rchg, nil
}

func upgradeVip(id uint32, openid, transid string) error {
	var tx *sql.Tx
	tx, err := DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET vip=1,points=points+520,rose=rose+520,date_privilege=date_privilege+1 WHERE open_id=? LIMIT 1`,
		openid,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	var rslt sql.Result
	rslt, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.recharge SET status=1,transaction_id=? WHERE id=? AND status=0 LIMIT 1`,
		transid, id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("upgradeVip: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("upgradeVip: recharge may not exists, or has already been handled")
	}
	return tx.Commit()
}

func rechargeRose(id, num uint32, openid, transid string) error {
	var tx *sql.Tx
	tx, err := DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET rose=rose+?,points=points+? WHERE open_id=? LIMIT 1`,
		num, num*10, openid,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	var rslt sql.Result
	rslt, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.recharge SET status=1,transaction_id=? WHERE id=? AND status=0 LIMIT 1`,
		transid, id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	if rslt == nil {
		tx.Rollback()
		return errors.New("rechargeRose: database driver error")
	}
	if affec, err := rslt.RowsAffected(); err != nil || affec != 1 {
		tx.Rollback()
		return errors.New("rechargeRose: recharge may not exists, or has already been handled")
	}
	return tx.Commit()
}
