/*
 * @Author: zhanghao
 * @DateTime: 2018-10-06 21:25:26
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 22:19:50
 */

package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/morgances/matchmaking/backend/util"
	"github.com/morgances/matchmaking/backend/conf"
)

type (
	userServPrvd struct{}

	User struct {
		// filled by member
		Phone            string
		Wechat           string
		NickName         string
		RealName         string
		Sex              uint8
		Birthday         string
		Height           string
		Location         string
		Job              string
		Faith            string
		SelfIntroduction string
		SelecCriteria    string

		OpenID        string
		Age           uint8
		CreateAt      time.Time
		Constellation string
		Certified     bool
		Vip           bool
		DatePrivilege uint32
		Points        float64
		Rose          uint32
		Charm         uint32
	}
)

var (
	UserService userServPrvd
)

func (userServPrvd) WeChatLogin(oid, nickName, loc string, sex uint8) error {
	exist, err := UserService.userExist(oid)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	return UserService.insert(&User{
		OpenID:   oid,
		NickName: nickName,
		Sex:      sex,
		Location: loc,
		CreateAt: time.Now(),
	})
}

func (userServPrvd) insert(u *User) error {
	_, err := DB.Exec(
		`INSERT INTO `+conf.MMConf.Database+`.user(open_id, nick_name, sex, location,create_at)
					VALUES(?,?,?,?,?,NOW())`,
		u.OpenID, u.NickName, u.Sex, u.Location,
	)
	return err
}

func (userServPrvd) userExist(oid string) (bool, error) {
	row := DB.QueryRow(
		`SELECT COUNT(0) FROM `+conf.MMConf.Database+`.user WHERE open_id = ? LOCK IN SHARE MODE`,
		oid,
	)
	var (
		err   error
		exist int32
	)
	if err = row.Scan(&exist); err != nil {
		return false, err
	}
	return exist == 1, nil
}

func (userServPrvd) FindByOpenID(oid string) (u *User, err error) {
	row := DB.QueryRow(
		`SELECT * FROM `+conf.MMConf.Database+`.user WHERE open_id = ? LOCK IN SHARE MODE`,
		oid,
	)

	u = &User{}
	if err = row.Scan(
		&u.Phone, &u.Wechat, &u.NickName, &u.RealName, &u.Sex, &u.Birthday, &u.Height,
		&u.Location, &u.Job, &u.Faith, &u.Constellation, &u.SelfIntroduction, &u.SelecCriteria,
		&u.OpenID, &u.Age, &u.CreateAt, &u.Certified, &u.Vip, &u.DatePrivilege, &u.Points, &u.Rose, &u.Charm,
	); err != nil {
		return nil, err
	}
	if len(u.Birthday) > 10 {
		u.Birthday = u.Birthday[:10]
	}
	return u, nil
}

func (userServPrvd) RecommendByCharm(sex uint8) (us []User, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM `+conf.MMConf.Database+`.user WHERE  sex=? ORDER BY charm DESC LOCK IN SHARE MODE`,
		sex,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		us = append(us, User{})
		err = rows.Scan(
			&us[i].Phone, &us[i].Wechat, &us[i].NickName, &us[i].RealName, &us[i].Sex, &us[i].Birthday, &us[i].Height,
			&us[i].Location, &us[i].Job, &us[i].Faith, &us[i].Constellation, &us[i].SelfIntroduction, &us[i].SelecCriteria,
			&us[i].OpenID, &us[i].Age, &us[i].CreateAt, &us[i].Certified, &us[i].Vip, &us[i].DatePrivilege, &us[i].Points, &us[i].Rose, &us[i].Charm,
		)
		if err != nil {
			return nil, err
		}
	}
	return us, nil
}

func (userServPrvd) GetContact(oid string) (phone, wechat string, err error) {
	row := DB.QueryRow(`SELECT phone, wechat FROM `+conf.MMConf.Database+`.user WHERE open_id=? LOCK IN SHARE MODE`, oid)
	if err = row.Scan(&phone, &wechat); err != nil {
		return "", "", err
	}
	return
}

func (userServPrvd) Certify(oid string) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET certified=1 WHERE open_id=? LIMIT 1`,
		oid,
	)
	return err
}

func (userServPrvd) DatePrivilegeReduce(oid string) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET date_privilege=date_privilege-1 WHERE open_id=? LIMIT 1`,
		oid,
	)
	return err
}

func (userServPrvd) DatePrivilegeAdd(oid string, num uint32) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET date_privilege=date_privilege+? WHERE open_id=? LIMIT 1`,
		num, oid,
	)
	return err
}

func (userServPrvd) Update(u *User) error {
	var err error
	u.Age, u.Constellation, err = util.GetAgeAndConstell(u.Birthday)
	if err != nil {
		return err
	}
	_, err = DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.user 
				  SET phone=?,wechat=?,nick_name=?,real_name=?,sex=?,birthday=?,height=?,location=?,job=?,faith=?,constellation=?,self_introduction=?,selec_criteria=?,
				  certified=?,vip=?,date_privilege=?,points=?,rose=?,charm=?,age=?
				  WHERE open_id=? LIMIT 1`,
		u.Phone, u.Wechat, u.NickName, u.RealName, u.Sex, u.Birthday, u.Height, u.Location, u.Job, u.Faith, u.Constellation, u.SelfIntroduction, u.SelecCriteria,
		u.Certified, u.Vip, u.DatePrivilege, u.Points, u.Rose, u.Charm, u.Age,
		u.OpenID,
	)
	return err
}

func (userServPrvd) SendRose(sender, recer string, num uint32) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET rose=rose-? WHERE open_id=? LIMIT 1`,
		num, sender,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("SendRose: rose not enough for (id): " + sender + " :" + err.Error())
	}
	_, err = tx.Exec(
		`UPDATE `+conf.MMConf.Database+`.user SET rose=rose+?, charm=charm+? WHERE open_id=? LIMIT 1`,
		num, num, recer,
	)
	if err != nil {
		tx.Rollback()
		return errors.New("SendRose: receiver error (id): " + sender + " :" + err.Error())
	}
	return tx.Commit()
}
