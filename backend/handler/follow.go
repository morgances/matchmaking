/*
 * Revision History:
 *     Initial: 2018/10/14        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/zh1014/comment/response"
	"github.com/dgrijalva/jwt-go"
)

type (
	shortUserInfo struct {
		OpenID   string `json:"open_id"`
		NickName string `json:"nick_name"`
	}
)

func Follow(this *server.Context) error {
	var (
		req    targetOpenID
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	err := this.JSONBody(&req)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if openid == req.TargetOpenID {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.FollowService.Insert(openid, req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

// Unfollow if exist
func Unfollow(this *server.Context) error {
	var (
		req targetOpenID
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	err := this.JSONBody(&req)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.FollowService.Delete(openid, req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetFollowing(this *server.Context) error {
	var (
		resp  struct {
			Following []shortUserInfo `json:"following"`
		}
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	users, err := model.FollowService.FindFollowing(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, user := range users {
		var shortUserInfo shortUserInfo
		shortUserInfo.OpenID = user.OpenID
		shortUserInfo.NickName = user.NickName
		resp.Following = append(resp.Following, shortUserInfo)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetFollower(this *server.Context) error {
	var (
		resp  []shortUserInfo
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	users, err := model.FollowService.FindFollower(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, user := range users {
		var shortUserInfo shortUserInfo
		shortUserInfo.OpenID = user.OpenID
		shortUserInfo.NickName = user.NickName
		resp = append(resp, shortUserInfo)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}
