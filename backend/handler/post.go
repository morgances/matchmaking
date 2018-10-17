/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/comment/response"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"strconv"
	"time"
)

type (
	post struct {
		ID      int64     `json:"id"`
		OpenID  string    `json:"open_id"`
		Title   string    `json:"title"`
		Content string    `json:"content"`
		Date    time.Time `json:"date"`
		Commend int64     `json:"commend"`
		Images  []string  `json:"Images"`
	}
)

func CreatePost(this *server.Context) error {
	var (
		err    error
		openid string
		postId int64
	)
	authorization := this.GetHeader("Authorization")
	openid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	title := this.FormValue("title")
	content := this.FormValue("content")
	post := &model.Post{
		OpenID:  openid,
		Title:   title,
		Content: content,
	}
	if postId, err = model.PostService.Insert(post); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}

	if err = util.SavePostImages(int(postId), this.Request()); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetReviewedPost(this *server.Context) error {
	var (
		err  error
		resp struct {
			Posts []post `json:"resp"`
		}
	)
	rawPosts, err := model.PostService.FindReviewed()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawPost := range rawPosts {
		post := post{}
		post.ID = rawPost.ID
		post.OpenID = rawPost.OpenID
		post.Title = rawPost.Title
		post.Content = rawPost.Content
		post.Date = rawPost.DateTime
		post.Commend = rawPost.Commend
		post.Images, _ = util.GetImages("./post/" + strconv.Itoa(int(post.ID)))
		resp.Posts = append(resp.Posts, post)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetMyPost(this *server.Context) error {
	var (
		err  error
		oid  string
		resp struct {
			Posts []post `json:"resp"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	rawPosts, err := model.PostService.FindByOpenID(oid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawPost := range rawPosts {
		post := post{}
		post.ID = rawPost.ID
		post.OpenID = rawPost.OpenID
		post.Title = rawPost.Title
		post.Content = rawPost.Content
		post.Date = rawPost.DateTime
		post.Commend = rawPost.Commend
		post.Images, _ = util.GetImages("./post/" + strconv.Itoa(int(post.ID)))
		resp.Posts = append(resp.Posts, post)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func CommendPost(this *server.Context) error {
	var (
		err error
		req targetID
	)
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.PostService.Commend(req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

// todo: delete image
func DeletePost(this *server.Context) error {
	var (
		err error
		oid string
		req targetID
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

	if err = model.PostService.DeleteByOpenIDAndID(oid, req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
