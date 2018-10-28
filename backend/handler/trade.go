/*
 * Revision History:
 *     Initial: 2018/10/14        Zhang Hao
 */

package handler

import (
	"time"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/zh1014/comment/response"
)

type (
	tradeInfo struct {
		ID        uint32    `json:"id"`
		OpenID    string    `json:"open_id"`
		GoodsID   uint32    `json:"goods_id"`
		BuyerName string    `json:"buyer_name"`
		GoodsName string    `json:"goods_name"`
		DateTime  time.Time `json:"date_time"`
		Cost      float64   `json:"cost"`
		Finished  bool      `json:"finished"`
	}
)

func CreateTrade(this *server.Context) error {
	var (
		err error
		req targetID
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
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

	trade := model.Trade{
		OpenID:  openid,
		GoodsID: req.TargetID,
	}
	u, err := model.UserService.FindByOpenID(openid)
	if err != nil {
		log.Error("CreateTrade: get buyer name by openid: ", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	trade.BuyerName = u.RealName
	g, err := model.GoodsService.FindByID(req.TargetID)
	if err != nil {
		log.Error("CreateTrade: get goods information by goodsid: ", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	trade.GoodsName = g.Title
	trade.Cost = g.Price
	if u.Vip {
		trade.Cost = trade.Cost * 0.88
	}

	if err = model.TradeService.Insert(&trade); err != nil {
		log.Error("CreateTrade: ", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetMyTrades(this *server.Context) error {
	var (
		resp []tradeInfo
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	rawTrades, err := model.TradeService.FindByOpenID(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rawTrade := range rawTrades {
		var tradeForFeed tradeInfo
		tradeForFeed.ID = rawTrade.ID
		tradeForFeed.OpenID = rawTrade.OpenID
		tradeForFeed.GoodsID = rawTrade.GoodsID
		tradeForFeed.BuyerName = rawTrade.BuyerName
		tradeForFeed.GoodsName = rawTrade.GoodsName
		tradeForFeed.Cost = rawTrade.Cost
		tradeForFeed.DateTime = rawTrade.DateTime
		tradeForFeed.Finished = rawTrade.Finished
		resp = append(resp, tradeForFeed)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}
