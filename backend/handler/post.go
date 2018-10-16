/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"log"
	"net/http"
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
		req    struct {
			Title    string `json:"title" validate:"required"`
			Content  string `json:"content"`
			ImageNum int    `json:"image_num" validate:"required, numeric, gte=0"`
		}
	)
	authorization := this.GetHeader("Authorization")
	openid, _, _, _, err = util.ParseToken(authorization)
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
	post := &model.Post{
		OpenID:  openid,
		Title:   req.Title,
		Content: req.Content,
	}
	if postId, err = model.PostService.Insert(post); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}

	if err = util.SavePostImages(req.ImageNum, int(postId), this.Request()); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrSaveImage)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetReviewedPost(this *server.Context) error {
	var (
		err  error
		resp struct {
			Posts []post `json:"resp"`
		}
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	rawPosts, err := model.PostService.FindReviewed()
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
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
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrLoadImage)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
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
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	rawPosts, err := model.PostService.FindByOpenID(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
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
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrLoadImage)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}

func CommendPost(this *server.Context) error {
	var (
		err error
		req targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.PostService.Commend(req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func DeletePost(this *server.Context) error {
	var (
		err error
		oid string
		req targetID
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.PostService.DeleteByOpenIDAndID(oid, req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}
