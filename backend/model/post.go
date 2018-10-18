/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 12:30:30
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:38:34
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type (
	postServPrvd struct{}

	Post struct {
		ID       int64
		OpenID   string
		Title    string
		Content  string
		DateTime time.Time
		Commend  int64
		Reviewed bool
	}
)

var (
	PostService postServPrvd
)

func (postServPrvd) Insert(p *Post) (int64, error) {
	result, err := DB.Exec(
		`INSERT INTO post(open_id,title,content,date_time)
					VALUES(?,?,?,NOW())
		`,
		p.OpenID, p.Title, p.Content,
	)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (postServPrvd) FindByID(id int64) (p *Post, err error) {
	row := DB.QueryRow(
		`SELECT * FROM post WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	p = &Post{}
	if err = row.Scan(
		&p.ID, &p.OpenID, &p.Title, &p.Content, &p.DateTime, &p.Commend, &p.Reviewed,
	); err != nil {
		return nil, err
	}
	return
}

func (postServPrvd) FindByOpenID(oid string) (ps []Post, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM post WHERE open_id=? ORDER BY date_time DESC LOCK IN SHARE MODE`,
		oid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Title, &ps[i].Content, &ps[i].DateTime, &ps[i].Commend, &ps[i].Reviewed,
		)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func (postServPrvd) FindUnreviewed() (ps []Post, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM post WHERE reviewed=0 ORDER BY date_time DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Title, &ps[i].Content, &ps[i].DateTime, &ps[i].Commend, &ps[i].Reviewed,
		)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func (postServPrvd) FindReviewed() (ps []Post, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM post WHERE reviewed=1 ORDER BY date_time DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		ps = append(ps, Post{})
		err = rows.Scan(
			&ps[i].ID, &ps[i].OpenID, &ps[i].Title, &ps[i].Content, &ps[i].DateTime, &ps[i].Commend, &ps[i].Reviewed,
		)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func (postServPrvd) UpdatePostStatus(id int64) error {
	rslt, err := DB.Exec(
		`UPDATE post SET reviewed=1 WHERE id=? LIMIT 1`,
		id,
	)
	if affected, err := rslt.RowsAffected(); err == nil && affected != 1 {
		return errors.New(fmt.Sprintf("failed to update status of post: %d", id))
	}
	return err
}

func (postServPrvd) Commend(id int64) error {
	rslt, err := DB.Exec(
		`UPDATE post SET commend=commend+1 WHERE id=? LIMIT 1`,
		id,
	)
	if affected, err := rslt.RowsAffected(); err == nil && affected != 1 {
		return errors.New(fmt.Sprintf("failed to commend post: %d", id))
	}
	return err
}

func (postServPrvd) DeleteByID(id int64) error {
	rslt, err := DB.Exec(
		`DELETE FROM post WHERE id=? LIMIT 1`,
		id,
	)
	if affected, err := rslt.RowsAffected(); err == nil && affected != 1 {
		return errors.New(fmt.Sprintf("failed to delete post: %d", id))
	}
	return err
}

func (postServPrvd) DeleteByOpenIDAndID(oid string, id int64) error {
	rslt, err := DB.Exec(
		`DELETE FROM post WHERE open_id=? AND id=? LIMIT 1`,
		oid, id,
	)
	if err != nil {
		return errors.New("DeleteByOpenIDAndID: " + err.Error())
	}
	if affected, err := rslt.RowsAffected(); err == nil && affected != 1 {
		return errors.New(fmt.Sprintf("user: %s failed to update status of post: %d", oid, id))
	}
	return err
}
