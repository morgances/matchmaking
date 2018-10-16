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
	"time"
)

type (
	tradeForResp struct {
		ID        int64     `json:"id"`
		OpenID    string    `json:"open_id"`
		GoodsID   int64     `json:"goods_id"`
		BuyerName string    `json:"buyer_name"`
		GoodsName string    `json:"goods_name"`
		DateTime  time.Time `json:"date_time"`
		Cost      int64     `json:"cost"`
		Finished  bool      `json:"finished"`
	}
)

func CreateTrade(this *server.Context) error {
	var (
		err error
		oid string
		req struct {
			GoodsID   int64  `json:"goods_id" validate:"required"`
			BuyerName string `json:"buyer_name" validate:"required"`
			GoodsName string `json:"goods_name" validate:"required"`
			Cost      int64  `json:"cost" validate:"required"`
		}
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

	trade := model.Trade{
		OpenID:    oid,
		GoodsID:   req.GoodsID,
		BuyerName: req.BuyerName,
		GoodsName: req.GoodsName,
		Cost:      req.Cost,
	}
	if err = model.TradeService.Insert(&trade); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetMyTrades(this *server.Context) error {
	var (
		err       error
		oid       string
		rawTrades []model.Trade
		resp      struct {
			Trades []tradeForResp `json:"trades"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	rawTrades, err = model.TradeService.FindByOpenID(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	for _, rawTrade := range rawTrades {
		var tradeForFeed tradeForResp
		tradeForFeed.ID = rawTrade.ID
		tradeForFeed.OpenID = rawTrade.OpenID
		tradeForFeed.GoodsID = rawTrade.GoodsID
		tradeForFeed.BuyerName = rawTrade.BuyerName
		tradeForFeed.GoodsName = rawTrade.GoodsName
		tradeForFeed.Cost = rawTrade.Cost
		tradeForFeed.DateTime = rawTrade.DateTime
		tradeForFeed.Finished = rawTrade.Finished
		resp.Trades = append(resp.Trades, tradeForFeed)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}
