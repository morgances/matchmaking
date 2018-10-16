/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
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
)

type (
	token struct {
		Token string `json:"token"`
	}

	targetID struct {
		TargetID int64 `json:"target_id"`
	}
)

func Login(this *server.Context) error {
	var (
		token token
	)
	authorization := this.GetHeader("Authorization")
	acc, pass, err := util.ParseBase64(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = model.AdminService.Login(acc, pass); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrAccountOrPasswordWrong)
	}
	if token.Token, err = util.NewToken("admin", "admin", 1, true); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.ServeJSON(&token); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}

func Certify(this *server.Context) error {
	var (
		err     error
		req     targetOpenID
		isAdmin bool
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
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
	if err = model.UserService.Certify(req.TargetOpenID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	err = this.WriteHeader(http.StatusOK)
	return err
}

func DatePrivilegeReduce(this *server.Context) error {
	var (
		err     error
		req     targetOpenID
		isAdmin bool
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	if err = model.UserService.DatePrivilegeReduce(req.TargetOpenID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func DatePrivilegeAdd(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetOpenID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	err = model.UserService.DatePrivilegeAdd(req.TargetOpenID, 1)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
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
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	resp.Phone, resp.Wechat, err = model.UserService.GetContact(req.TargetOpenID)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	if err = this.ServeJSON(resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetUnreviewedPost(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		resp    struct {
			Posts []post `json:"resp"`
		}
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	rawPosts, err := model.PostService.FindUnreviewed()
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
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}

func UpdatePostStatus(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.PostService.UpdatePostStatus(req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func AdminDeletePost(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.PostService.DeleteByID(req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
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
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	rawTrades, err := model.TradeService.FindUnfinishedTrade()
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
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
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}

func CancelTrade(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     struct {
					ID     int64  `json:"id"`
					OpenID string `json:"open_id"`
					Cost   int64  `json:"cost"`
				}
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	trade := model.Trade{
		ID:     req.ID,
		OpenID: req.OpenID,
		Cost:   req.Cost,
	}
	if err = model.TradeService.Cancel(&trade); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func UpdateTradeStatus(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.TradeService.UpdateTradeStatus(req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}