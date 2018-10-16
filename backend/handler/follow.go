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

type (
	shortUserInfo struct {
		OpenID   string `json:"open_id"`
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
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
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	if err = model.FollowService.Insert(oid, req.TargetOpenID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func Unfollow(this *server.Context) error {
	var (
		err error
		oid string
		req targetOpenID
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	if err = model.FollowService.Delete(oid, req.TargetOpenID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
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
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	users, err = model.FollowService.FindFollowing(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	for _, user := range users {
		var shortUserInfo shortUserInfo
		shortUserInfo.OpenID = user.OpenID
		shortUserInfo.NickName = user.NickName
		shortUserInfo.Avatar = user.Avatar
		resp.Following = append(resp.Following, shortUserInfo)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
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
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	users, err = model.FollowService.FindFollower(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	for _, user := range users {
		var shortUserInfo shortUserInfo
		shortUserInfo.OpenID = user.OpenID
		shortUserInfo.NickName = user.NickName
		shortUserInfo.Avatar = user.Avatar
		resp.Follower = append(resp.Follower, shortUserInfo)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}
