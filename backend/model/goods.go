/*
 * @Author: zhanghao
 * @Date: 2018-10-08 22:46:12
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:38:07
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type (
	goodsServPrvd struct{}

	Goods struct {
		ID          int64
		Name        string
		Price       int64
		Description string
		Image       string
	}
)

var (
	GoodsService goodsServPrvd
)

func (goodsServPrvd) Insert(g *Goods) error {
	_, err := DB.Exec(
		`INSERT INTO goods(id,name,price,description,image)
			VALUES(?,?,?,?,?)
		`,
		g.ID, g.Name, g.Price, g.Description, g.Image,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New(fmt.Sprintf("duplicate entry id:%s", g.ID)) // need fix when struct field changed
		}
		return err
	}
	return nil
}

func (goodsServPrvd) FindByID(id int64) (g *Goods, err error) {
	row := DB.QueryRow(
		`SELECT * FROM goods WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	if err = row.Scan(
		&g.ID, &g.Name, &g.Price, &g.Description, &g.Image,
	); err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return nil, err
}

func (goodsServPrvd) FindByTime() (gs []Goods, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(
		`SELECT * FROM goods ORDER BY price DESC LOCK IN SHARE MODE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		gs = append(gs, Goods{})
		err = rows.Scan(
			&gs[i].ID, &gs[i].Name, &gs[i].Price, &gs[i].Description, &gs[i].Image,
		)
		if err != nil {
			return nil, err
		}
	}
	return gs, nil
}

func (goodsServPrvd) Update(g *Goods) error {
	_, err := GoodsService.FindByID(g.ID)
	if err != nil {
		return err
	}
	_, err = DB.Exec(
		`UPDATE goods SET
			id=?,name=?,price=?,description=?,image=?
			WHERE id=? LIMIT 1
		`,
		g.ID, g.Name, g.Price, g.Description, g.Image,
		g.ID,
	)
	return err
}

func (goodsServPrvd) DeleteByID(id int64) error {
	_, err := DB.Exec(
		`DELETE FROM goods WHERE id=? LIMIT 1`,
		id,
	)
	return err
}
