/*
 * Revision History:
 *     Initial: 2018/09/12        Tong Yuehong
 */

package controller

import (
	"database/sql"

	"github.com/TechCatsLab/apix/http/server"
	"github.com/zh1014/comment/config"
	"github.com/zh1014/comment/constants"
	"github.com/zh1014/comment/response"
	"github.com/zh1014/comment/services"
	log "github.com/TechCatsLab/logging/logrus"
)

type Controller struct {
	service *services.CommentService
}

func New(db *sql.DB, c *config.Config) *Controller {
	return &Controller{
		service: services.NewService(c, db),
	}
}

func (con *Controller) CreateDB() error {
	return con.service.CreateDB()
}

func (con *Controller) CreateTable() error {
	return con.service.CreateTable()
}

func (con *Controller) Insert(c *server.Context) error {
	var (
		req struct {
			TargetID uint64 `json:"target_id"`
			Content  string `json:"content"`
			UserID   string `json:"user_id"`
			ParentID uint64 `json:"parent_id"`
		}
	)

	if err := c.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	id, err := con.service.Insert(req.TargetID, req.ParentID, req.UserID, req.Content)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrMysql, nil)
	}

	return response.WriteStatusAndIDJSON(c, constants.ErrSucceed, id)
}

func (con *Controller) ChangeStatus(c *server.Context) error {
	var (
		req struct {
			CommentID uint64 `json:"comment_id"`
			Status    uint8  `json:"status"`
		}
	)

	if err := c.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	err := con.service.ChangeStatus(req.CommentID, req.Status)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrMysql, nil)
	}

	return response.WriteStatusAndDataJSON(c, constants.ErrSucceed, nil)
}

func (con *Controller) ChangeContent(c *server.Context) error {
	var (
		req struct {
			CommentID uint64 `json:"comment_id"`
			Content   string `json:"content"`
		}
	)

	if err := c.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	err := con.service.ChangeContent(req.CommentID, req.Content)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrMysql, nil)
	}

	return response.WriteStatusAndDataJSON(c, constants.ErrSucceed, nil)
}

func (con *Controller) ListCommentsByTarget(c *server.Context) error {
	var (
		req struct {
			TargetID uint64 `json:"target_id"`
		}
	)

	if err := c.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	comments, err := con.service.CommentsByTarget(req.TargetID)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrMysql, nil)
	}

	return response.WriteStatusAndDataJSON(c, constants.ErrSucceed, comments)
}

func (con *Controller) ListCommentsByUserID(c *server.Context) error {
	var (
		req struct {
			UserID string `json:"user_id"`
		}
	)

	if err := c.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	comments, err := con.service.CommentsByUser(req.UserID)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(c, constants.ErrMysql, nil)
	}

	return response.WriteStatusAndDataJSON(c, constants.ErrSucceed, comments)
}
