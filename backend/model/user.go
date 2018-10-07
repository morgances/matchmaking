/*
 * @Author: zhanghao
 * @Date: 2018-10-06 21:25:26
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-07 18:26:41
 */

package model

import (
	"database/sql"
	"time"
)

type (
	UserServPrvd struct{}

	User struct {
		// filled by member
		Phone            string
		Wechat           string
		NickName         string
		Avatar           string
		RealName         string
		Sex              bool
		Birthday         time.Time
		Height           string
		Location         string
		Job              string
		Faith            string
		Constellation    string
		SelfIntroduction string
		SelecCriteria    string

		CreateAt      time.Time
		Password      string
		Album         string
		Certified     bool
		Vip           bool
		DatePrivilege int64
		Points        int64
		Rose          int64
		Charm         int64
	}
)

var UserService UserServPrvd

func (UserServPrvd) Insert(u *User) error {
	_, err := DB.Exec(
		`INSERT INTO user(phone,wechat,nick_name,avatar,real_name,sex,birthday,height,location,job,faith,constellation,self_introduction,selec_criteria,
			create_at,password,album,certified,vip,date_privilege,points,rose,charm)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,NOW(),?,?,?,?,?,?,?,?);`,
		u.Phone, u.Wechat, u.NickName, u.Avatar, u.RealName, u.Sex, u.Birthday, u.Height, u.Location, u.Job, u.Faith, u.Constellation, u.SelfIntroduction, u.SelecCriteria,
		u.Password, u.Album, u.Certified, u.Vip, u.DatePrivilege, u.Points, u.Rose, u.Charm,
	)
	return err
}

func (UserServPrvd) FindByPhone(p string) (u User, err error) {
	row := DB.QueryRow(
		`SELECT * FROM user WHERE phone = ?`,
		p,
	)
	err = row.Scan(
		&u.Phone, &u.Wechat, &u.NickName, &u.Avatar, &u.RealName, &u.Sex, &u.Birthday, &u.Height,
		&u.Location, &u.Job, &u.Faith, &u.Constellation, &u.SelfIntroduction, &u.SelecCriteria,
		&u.CreateAt, &u.Password, &u.Album, &u.Certified, &u.Vip, &u.DatePrivilege, &u.Points, &u.Rose, &u.Charm,
	)
	return
}

func (UserServPrvd) RecommendByCharm() (us []User, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM user ORDER BY charm DESC`,
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
	us = make([]User, len(cols))
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(
			&us[i].Phone, &us[i].Wechat, &us[i].NickName, &us[i].Avatar, &us[i].RealName, &us[i].Sex, &us[i].Birthday, &us[i].Height,
			&us[i].Location, &us[i].Job, &us[i].Faith, &us[i].Constellation, &us[i].SelfIntroduction, &us[i].SelecCriteria,
			&us[i].CreateAt, &us[i].Password, &us[i].Album, &us[i].Certified, &us[i].Vip, &us[i].DatePrivilege, &us[i].Points, &us[i].Rose, &us[i].Charm,
		)
		if err != nil {
			return nil, err
		}
	}
	return us, nil
}

func (UserServPrvd) Update(u *User) error {
	_, err := DB.Exec(
		`UPDATE user SET 
		phone=?,wechat=?,nick_name=?,avatar=?,real_name=?,sex=?,birthday=?,height=?,location=?,job=?,faith=?,constellation=?,self_introduction=?,selec_criteria=?,
		create_at=?,password=?,album=?,certified=?,vip=?,date_privilege=?,points=?,rose=?,charm=?
		WHERE phone=?;`,
		u.Phone, u.Wechat, u.NickName, u.Avatar, u.RealName, u.Sex, u.Birthday, u.Height, u.Location, u.Job, u.Faith, u.Constellation, u.SelfIntroduction, u.SelecCriteria,
		u.CreateAt, u.Password, u.Album, u.Certified, u.Vip, u.DatePrivilege, u.Points, u.Rose, u.Charm,
		u.Phone,
	)
	return err
}

func (UserServPrvd) DeleteByPhone(p string) error {
	_, err := DB.Exec(
		`DELETE FROM user WHERE phone=?`,
		p,
	)
	return err
}
