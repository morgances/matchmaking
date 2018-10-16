/*
 * Revision History:
 *     Initial: 2018/10/14        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"log"
	"net/http"
)

func Signin(this *server.Context) error {
	var (
		err error
		oid string
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	if err = model.SigninService.Insert(oid); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
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
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	resp.SigninRecord, err = model.SigninService.FindByOpenID(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}
