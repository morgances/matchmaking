/*
 * Revision History:
 *     Initial: 2018/10/14        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/comment/response"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
)

type (
	shortUserInfo struct {
		OpenID   string `json:"open_id"`
		NickName string `json:"nick_name"`
	}
)

func Follow(this *server.Context) error {
	var (
		err error
		oid string
		req targetOpenID
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if oid == req.TargetOpenID {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.FollowService.Insert(oid, req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

// Unfollow if exist
func Unfollow(this *server.Context) error {
	var (
		err error
		oid string
		req targetOpenID
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.FollowService.Delete(oid, req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetFollowing(this *server.Context) error {
	var (
		err   error
		oid   string
		users []model.User
		resp  struct {
			Following []shortUserInfo `json:"following"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	users, err = model.FollowService.FindFollowing(oid)
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
		err   error
		oid   string
		users []model.User
		resp  struct {
			Follower []shortUserInfo `json:"follower"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	users, err = model.FollowService.FindFollower(oid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, user := range users {
		var shortUserInfo shortUserInfo
		shortUserInfo.OpenID = user.OpenID
		shortUserInfo.NickName = user.NickName
		resp.Follower = append(resp.Follower, shortUserInfo)
	}
	log.Error(err)
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}
