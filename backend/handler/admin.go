/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
 */

package handler

import (
	"strconv"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"github.com/morgances/matchmaking/backend/wx"
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
		resp token
	)
	authorization := this.GetHeader("Authorization")
	acc, pass, err := wx.ParseBase64(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = model.AdminService.Login(acc, pass); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrAccount, nil)
	}
	if resp.Token, err = wx.NewToken("admin", 1, true); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func Certify(this *server.Context) error {
	var (
		err     error
		req     targetOpenID
		isAdmin bool
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	//////todo:
	//token := this.Request().Context().Value("is_admin").(*jwt.Token)
	//isAdmin = token.Header["is_admin"].(bool)
	//println(isAdmin)
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		req     targetOpenID
		isAdmin bool
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		isAdmin bool
		req     targetOpenID
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		isAdmin bool
		req     targetOpenID
		resp    struct {
			Phone  string `json:"phone"`
			Wechat string `json:"wechat"`
		}
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		isAdmin bool
		resp    struct {
			Posts []post `json:"posts"`
		}
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	rawPosts, err := model.PostService.FindUnreviewed()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
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

func UpdatePostStatus(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
	if err = util.ClearPostImages(req.TargetID); err != nil {
		// make a log but tell admin succeed, because it succeed in database
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetUnfinishedTrade(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		resp    struct {
			Trades []tradeForResp `json:"trades"`
		}
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}

	rawTrades, err := model.TradeService.FindUnfinishedTrade()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawTrade := range rawTrades {
		tradeForFeed := tradeForResp{}
		tradeForFeed.ID = rawTrade.ID
		tradeForFeed.OpenID = rawTrade.OpenID
		tradeForFeed.GoodsID = rawTrade.GoodsID
		tradeForFeed.DateTime = rawTrade.DateTime
		tradeForFeed.Cost = rawTrade.Cost
		resp.Trades = append(resp.Trades, tradeForFeed)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func CancelTrade(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, isAdmin, err = wx.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
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
