/*
 * Revision History:
 *     Initial: 2018/10/19        Zhang Hao
 */

package wx

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"

	"github.com/193Eric/go-wechat"
)

type OrderInfo struct {
	AppID          string
	Body           string
	MchID          string
	NonceStr       string
	SpbillCreateIP string
	TotalFee       int
	OutTradeNo     string
	NotifyUrl      string
	TradeType      string
	OpenID         string
	Key            string
}

func VipOrderInfo(spbillCreateIP, outTradeNo, openID string) *OrderInfo {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &OrderInfo{
		AppID:     conf.MMConf.AppID,
		Body:      constant.RechargeVIPBody,
		MchID:     conf.MMConf.MchID,
		TotalFee:  conf.MMConf.VIPFee,
		NotifyUrl: "https://" + conf.MMConf.Address + ":" + conf.MMConf.Port + constant.NotifyUrl,
		TradeType: constant.TradeType,
		Key:       conf.MMConf.AppOrderKey,

		NonceStr:       fmt.Sprintf("%d%s", r.Intn(10000), openID[5:]),
		SpbillCreateIP: spbillCreateIP,
		OutTradeNo:     outTradeNo,
		OpenID:         openID,
	}
}

func RoseOrder(spbillCreateIP, outTradeNo, openID string, num int) *OrderInfo {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &OrderInfo{
		AppID:     conf.MMConf.AppID,
		Body:      constant.RechargeRoseBody,
		MchID:     conf.MMConf.MchID,
		TotalFee:  num * 100,
		NotifyUrl: "https://" + conf.MMConf.Address + ":" + conf.MMConf.Port + constant.NotifyUrl,
		TradeType: constant.TradeType,
		Key:       conf.MMConf.AppOrderKey,

		NonceStr:       fmt.Sprintf("%d%s", r.Intn(10000), openID[5:]),
		SpbillCreateIP: spbillCreateIP,
		OutTradeNo:     outTradeNo,
		OpenID:         openID,
	}
}

func SetOrder(orderinfo *OrderInfo) (*wechat.UnifyOrderResp, error) {
	return wechat.SetOrder(orderinfo.AppID, orderinfo.Body, orderinfo.MchID, orderinfo.NonceStr, orderinfo.SpbillCreateIP,
		orderinfo.TotalFee, orderinfo.OutTradeNo, orderinfo.NotifyUrl, orderinfo.TradeType, orderinfo.OpenID, orderinfo.Key)
}

// f(Out_trade_no,Transaction_id,Result_code)
func PayCallback(ctx *server.Context, f func(string, string, string)) (string, string) {
	return wechat.WxpayCallback(ctx.Response(), ctx.Request(), f, conf.MMConf.AppOrderKey)
}

func HandleRecharge(outTradeNo, transactionID, resultCode string) {
	outNum, err := strconv.Atoi(outTradeNo)
	if err != nil {
		log.Error("HandleRecharge convert out_trade_no to number: " + err.Error())
	}
	switch resultCode {
	case "SUCCESS":
		err = model.RechargeService.Success(outNum, transactionID)
		if err != nil {
			log.Error("HandleRecharge: ", err)
		}
	case "FAIL":
		err = model.RechargeService.Fail(outNum)
		if err != nil {
			log.Error("HandleRecharge: ", err)
		}
	default:
		log.Error("HandleRecharge unknown resultCode " + resultCode)
	}
}
