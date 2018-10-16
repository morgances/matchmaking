/*
 * Revision History:
 *     Initial: 2018/09/12        Tong Yuehong
 */

package main

import (
	"database/sql"
	"errors"

	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/comment/config"
	"github.com/TechCatsLab/comment/controller"
)

var (
	errInvalidRouter = errors.New("[RegisterRouter]: server is nil")
)

func Register(r *server.Router, db *sql.DB, c *config.Config) error {
	if r == nil {
		return errInvalidRouter
	}

	con := controller.New(db, c)
	err := con.CreateDB()
	if err != nil {
		return err
	}

	err = con.CreateTable()
	if err != nil {
		return err
	}

	r.Post("/api/v1/"+c.RequestDomain+"/comment/insert", con.Insert)
	r.Post("/api/v1/"+c.RequestDomain+"/comment/changestatus", con.ChangeStatus)
	r.Post("/api/v1/"+c.RequestDomain+"/comment/changecontent", con.ChangeContent)
	r.Post("/api/v1/"+c.RequestDomain+"/comment/target", con.ListCommentsByTarget)
	r.Post("/api/v1/"+c.RequestDomain+"/comment/user", con.ListCommentsByUserID)

	return nil
}
