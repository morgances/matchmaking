/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 12:30:30
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:38:34
 */

package model

import (
	"database/sql"
	"time"

	"github.com/morgances/matchmaking/backend/conf"
)

type (
	postServPrvd struct{}

	Post struct {
		ID       uint32
		OpenID   string
		Content  string
		DateTime time.Time
		Commend  uint32
		Reviewed bool

		NickName      string
		VIP           bool
		Age           uint8
		Location      string
		Height        string
		Constellation string
	}
)

var (
	PostService postServPrvd

	// Users look through the newest post
	unrvwdSQL = `SELECT p.id,p.open_id,p.content,p.date_time,p.commend,u.nick_name,u.vip,u.age,u.location,u.height,u.constellation 
					FROM `+conf.MMConf.Database+`.post p 
					JOIN `+conf.MMConf.Database+`.user u ON p.open_id=u.open_id 
					WHERE p.reviewed=0 
					ORDER BY date_time ASC LOCK IN SHARE MODE`
	// admin looks through the oldest post
	rvwdSQL = `SELECT p.id,p.open_id,p.content,p.date_time,p.commend,u.nick_name,u.vip,u.age,u.location,u.height,u.constellation 
					FROM `+conf.MMConf.Database+`.post p 
					JOIN `+conf.MMConf.Database+`.user u ON p.open_id=u.open_id 
					WHERE p.reviewed=1 
					ORDER BY date_time DESC LOCK IN SHARE MODE`
)

func (postServPrvd) Insert(p *Post) (uint32, error) {
	result, err := DB.Exec(
		`INSERT INTO `+conf.MMConf.Database+`.post(open_id,content,date_time)
					VALUES(?,?,NOW())
		`,
		p.OpenID, p.Content,
	)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(lastId), nil
}

func (postServPrvd) FindByOpenID(oid string) (ps []Post, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT p.id,p.open_id,p.content,p.date_time,p.commend 
					FROM `+conf.MMConf.Database+`.post 
					WHERE open_id=? 
					ORDER BY date_time DESC LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Content, &ps[i].DateTime, &ps[i].Commend,
		)
		if err != nil {
			return nil, err
		}
	}
	return ps, rows.Err()
}

func (postServPrvd) FindMany(isreviewed bool) (ps []Post, err error) {
	var (
		rows  *sql.Rows
		SQL string
	)

	if isreviewed {
		SQL = rvwdSQL
	} else {
		SQL = unrvwdSQL
	}

	if rows, err = DB.Query(SQL); err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		if err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Content, &ps[i].DateTime, &ps[i].Commend, &ps[i].NickName, &ps[i].VIP, &ps[i].Age, &ps[i].Location, &ps[i].Height, &ps[i].Constellation,
		); err != nil {
			return nil, err
		}
	}
	return ps, rows.Err()
}

func (postServPrvd) UpdatePostStatus(id uint32) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.post SET reviewed=1 WHERE id=? LIMIT 1`,
		id,
	)
	return err
}

func (postServPrvd) Commend(id uint32) error {
	_, err := DB.Exec(
		`UPDATE `+conf.MMConf.Database+`.post SET commend=commend+1 WHERE id=? LIMIT 1`,
		id,
	)
	return err
}

func (postServPrvd) DeleteByID(id uint32) error {
	_, err := DB.Exec(
		`DELETE FROM `+conf.MMConf.Database+`.post WHERE id=? LIMIT 1`,
		id,
	)
	return err
}

func (postServPrvd) DeleteByOpenIDAndID(oid string, id uint32) error {
	_, err := DB.Exec(
		`DELETE FROM `+conf.MMConf.Database+`.post WHERE open_id=? AND id=? LIMIT 1`,
		oid, id,
	)
	return err
}
