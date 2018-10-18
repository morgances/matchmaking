/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 22:46:12
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:38:07
 */

package model

import (
	"database/sql"
	"errors"
	"fmt"
)

type (
	goodsServPrvd struct{}

	Goods struct {
		ID          int64
		Title       string
		Price       float64
		Description string
	}
)

var (
	GoodsService goodsServPrvd
)

func (goodsServPrvd) Insert(g *Goods) (int64, error) {
	result, err := DB.Exec(
		`INSERT INTO goods(title,price,description)
					VALUES(?,?,?)`,
		g.Title, g.Price, g.Description,
	)
	if err != nil {
		return 0, err
	}
	var lastId int64
	lastId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, err
}

func (goodsServPrvd) FindByID(id int64) (g *Goods, err error) {
	row := DB.QueryRow(
		`SELECT * FROM goods WHERE id=? LOCK IN SHARE MODE`,
		id,
	)
	g = &Goods{}
	if err = row.Scan(
		&g.ID, &g.Title, &g.Price, &g.Description,
	); err != nil {
		return nil, err
	}
	return g, nil
}

func (goodsServPrvd) FindByPrice() (gs []Goods, err error) {
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
			&gs[i].ID, &gs[i].Title, &gs[i].Price, &gs[i].Description,
		)
		if err != nil {
			return nil, err
		}
	}
	return gs, nil
}

func (goodsServPrvd) Update(g *Goods) error {
	_, err := DB.Exec(
		`UPDATE goods 
					SET title=?,price=?,description=?
					WHERE id=? LIMIT 1`,
		g.Title, g.Price, g.Description,
		g.ID,
	)
	return err
}

func (goodsServPrvd) DeleteByID(id int64) error {
	var rslt sql.Result
	rslt, err := DB.Exec(
		`DELETE FROM goods WHERE id=? LIMIT 1`,
		id,
	)
	if err != nil {
		return err
	}
	if affected, err := rslt.RowsAffected(); err == nil && affected != 1 {
		return errors.New(fmt.Sprintf("failed to delete good: %d", id))
	}
	return err
}
