/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
 */

package handler

import (
	"github.com/morgances/matchmaking/backend/img"

	"github.com/dgrijalva/jwt-go"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"github.com/zh1014/comment/response"
)

type (
	token struct {
		Token string `json:"token"`
	}

	targetID struct {
		TargetID uint32 `json:"target_id" validate:"required,gte=1"`
	}
)

func Login(this *server.Context) error {
	var (
		req struct {
			Account  string `json:"admin_account" validate:"required"`
			Password string `json:"admin_password" validate:"required"`
		}
		resp token
	)
	err := this.JSONBody(&req)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = model.AdminService.Login(req.Account, req.Password); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrAccount, nil)
	}
	if resp.Token, err = util.NewToken("admin", 1, true); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func Certify(this *server.Context) error {
	var (
		req targetOpenID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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
	if err = model.UserService.Certify(req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func DatePrivilegeReduce(this *server.Context) error {
	var (
		req targetOpenID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	if err = model.UserService.DatePrivilegeReduce(req.TargetOpenID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func DatePrivilegeAdd(this *server.Context) error {
	var (
		req targetOpenID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	err = model.UserService.DatePrivilegeAdd(req.TargetOpenID, 1)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetContact(this *server.Context) error {
	var (
		req  targetOpenID
		resp struct {
			Phone  string `json:"phone"`
			Wechat string `json:"wechat"`
		}
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	resp.Phone, resp.Wechat, err = model.UserService.GetContact(req.TargetOpenID)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetUnreviewedPost(this *server.Context) error {
	var (
		resp []post
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	rawPosts, err := model.PostService.FindMany(false)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
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
		post.Images = img.GetPostImgs(post.ID)
		resp = append(resp, post)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func UpdatePostStatus(this *server.Context) error {
	var (
		req targetID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	if err = model.PostService.UpdatePostStatus(req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func AdminDeletePost(this *server.Context) error {
	var (
		req targetID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	if err = model.PostService.DeleteByID(req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	if err = img.ClearPostImages(req.TargetID); err != nil {
		// make a log but tell admin succeed, because it succeed in database
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetUnfinishedTrade(this *server.Context) error {
	var (
		resp []tradeInfo
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}

	rawTrades, err := model.TradeService.FindUnfinishedTrade()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawTrade := range rawTrades {
		tradeForFeed := tradeInfo{}
		tradeForFeed.ID = rawTrade.ID
		tradeForFeed.OpenID = rawTrade.OpenID
		tradeForFeed.GoodsID = rawTrade.GoodsID
		tradeForFeed.DateTime = rawTrade.DateTime
		tradeForFeed.Cost = rawTrade.Cost
		resp = append(resp, tradeForFeed)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func CancelTrade(this *server.Context) error {
	var (
		req targetID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	tradeRecord, err := model.TradeService.FindByID(req.TargetID)
	if err != nil {
		log.Error("CancelTrade: ", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	trade := model.Trade{
		ID:     tradeRecord.ID,
		OpenID: tradeRecord.OpenID,
		Cost:   tradeRecord.Cost,
	}
	if err = model.TradeService.Cancel(&trade); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func UpdateTradeStatus(this *server.Context) error {
	var (
		req targetID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
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

	if err = model.TradeService.UpdateTradeStatus(req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
