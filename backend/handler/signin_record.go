/*
 * Revision History:
 *     Initial: 2018/10/14        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/zh1014/comment/response"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/wx"
)

func Signin(this *server.Context) error {
	var (
		err error
		oid string
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	if err = model.SigninService.Insert(oid); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetSigninRecord(this *server.Context) error {
	var (
		err  error
		oid  string
		resp struct {
			SigninRecord []string `json:"signin_record"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	resp.SigninRecord, err = model.SigninService.FindByOpenID(oid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}
