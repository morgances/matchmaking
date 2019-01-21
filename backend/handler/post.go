/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"strconv"
	"time"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/img"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/zh1014/comment/response"
)

type (
	post struct {
		ID      uint32    `json:"id"`
		OpenID  string    `json:"open_id"`
		Content string    `json:"content"`
		Date    time.Time `json:"date"`
		Commend uint32    `json:"commend"`
		Images  []string  `json:"Images"`

		NickName      string `json:"nick_name"`
		VIP           bool   `json:"vip"`
		Age           uint8  `json:"age"`
		Location      string `json:"location"`
		Height        string `json:"height"`
		Constellation string `json:"constellation"`
	}
)

func CreatePost(this *server.Context) error {
	// req form-data: image_num title content
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	content := this.FormValue("content")
	post := &model.Post{
		OpenID:  openid,
		Content: content,
	}
	postId, err := model.PostService.Insert(post)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}

	if err = img.SavePostImages(postId, this.Request()); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetReviewedPost(this *server.Context) error {
	var (
		err  error
		resp []post
	)
	rawPosts, err := model.PostService.FindMany(true)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawPost := range rawPosts {
		post := post{}
		post.ID = rawPost.ID
		post.OpenID = rawPost.OpenID
		post.Content = rawPost.Content
		post.Date = rawPost.DateTime
		post.Commend = rawPost.Commend
		post.NickName = rawPost.NickName
		post.VIP = rawPost.VIP
		post.Age = rawPost.Age
		post.Location = rawPost.Location
		post.Height = rawPost.Height
		post.Constellation = rawPost.Constellation
		post.Images, _ = img.GetImages("./post/" + strconv.Itoa(int(post.ID)))
		resp = append(resp, post)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetMyPost(this *server.Context) error {
	var (
		resp []post
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	rawPosts, err := model.PostService.FindByOpenID(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawPost := range rawPosts {
		post := post{}
		post.ID = rawPost.ID
		post.OpenID = rawPost.OpenID
		post.Content = rawPost.Content
		post.Date = rawPost.DateTime
		post.Commend = rawPost.Commend
		post.Images, _ = img.GetImages("./post/" + strconv.Itoa(int(post.ID)))
		resp = append(resp, post)
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

func DeletePost(this *server.Context) error {
	var (
		req targetID
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

	if err = model.PostService.DeleteByOpenIDAndID(openid, req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	if err = img.ClearPostImages(req.TargetID); err != nil {
		// make a log but tell user succeed, because it succeed in database
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
