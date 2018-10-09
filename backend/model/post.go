/*
 * @Author: zhanghao
 * @Date: 2018-10-08 12:30:30
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:38:34
 */

package model

import (
	"database/sql"
	"time"
)

type (
	postServPrvd struct{}

	Post struct {
		ID      int64
		OpenID  string
		Title   string
		Image   string
		Content string
		Date    time.Time
		Like    int64
	}
)

var (
	PostService postServPrvd
)

func (postServPrvd) Insert(p *Post) error {
	_, err := DB.Exec(
		`INSERT INTO post(open_id,title,image,content,date_time,like)
			VALUES(?,?,?,?,?,?)
		`,
		p.OpenID, p.Title, p.Image, p.Content, p.Date, p.Like,
	)

	if err != nil {
		return err
	}
	return nil
}

func (postServPrvd) FindByID(id int64) (p *Post, err error) {
	row := DB.QueryRow(
		`SELECT * FROM post WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	if err = row.Scan(
		&p.ID, &p.OpenID, &p.Title, &p.Image, &p.Content, &p.Date, &p.Like,
	); err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return nil, err
}

func (postServPrvd) FindByTime() (ps []Post, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM post ORDER BY date_time DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Title, &ps[i].Image, &ps[i].Content, &ps[i].Date, &ps[i].Like,
		)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func (postServPrvd) Update(p *Post) error {
	_, err := DB.Exec(
		`UPDATE post SET 
			open_id=?,title=?,image=?,content=?,date_time=?,like=?
			WHERE id=? LIMIT 1
		`,
		p.OpenID, p.Title, p.Image, p.Content, p.Date, p.Like,
		p.ID,
	)
	return err
}

func (postServPrvd) DeleteByID(id int64) error {
	_, err := DB.Exec(
		`DELETE FROM post WHERE id=? LIMIT 1`,
		id,
	)
	return err
}

func (postServPrvd) DeleteSomeoneAllPost(oid string) error {
	_, err := DB.Exec(
		`DELETE FROM post WHERE open_id=?`,
		oid,
	)
	return err
}
