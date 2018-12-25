/*
 * Revision History:
 *     Initial: 2018/10/19        Zhang Hao
 */

package wx

import (
	"strconv"
	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"

	"github.com/193Eric/go-wechat"
	"sort"
	"fmt"
	"crypto/md5"
	"strings"
	"encoding/hex"
)

var (
	// TODO: use https
	notifyURL = "http://" + conf.MMConf.Address + ":" + conf.MMConf.Port + constant.NotifyUrl
)

type OrderInfo struct {
	AppID          string
	Body           string
	MchID          string
	NonceStr       string
	SpbillCreateIP string
	TotalFee       uint32
	OutTradeNo     string
	NotifyUrl      string
	TradeType      string
	OpenID         string
	Key            string
}

// VipOrderInfo create a OrderInfo instance for recharge vip
func VipOrderInfo(spbillCreateIP, outTradeNo, openID string) *OrderInfo {
	return &OrderInfo{
		AppID:     conf.MMConf.AppID,
		Body:      constant.RechargeVIPBody,
		MchID:     conf.MMConf.MchID,
		TotalFee:  conf.MMConf.VIPFee,
		NotifyUrl: notifyURL,
		TradeType: constant.TradeType,
		Key:       conf.MMConf.AppOrderKey,

		NonceStr:       util.RandomStr(20),
		SpbillCreateIP: spbillCreateIP,
		OutTradeNo:     outTradeNo,
		OpenID:         openID,
	}
}

// RoseOrder create a OrderInfo instance for recharge rose
func RoseOrder(spbillCreateIP, outTradeNo, openID string, num uint32) *OrderInfo {
	return &OrderInfo{
		AppID:     conf.MMConf.AppID,
		Body:      constant.RechargeRoseBody,
		MchID:     conf.MMConf.MchID,
		TotalFee:  num * 100,
		NotifyUrl: notifyURL,
		TradeType: constant.TradeType,
		Key:       conf.MMConf.AppOrderKey,

		NonceStr:       util.RandomStr(20),
		SpbillCreateIP: spbillCreateIP,
		OutTradeNo:     outTradeNo,
		OpenID:         openID,
	}
}

// SetOrder send a request to wechat for creating order
func SetOrder(orderinfo *OrderInfo) (*wechat.UnifyOrderResp, error) {
	return wechat.SetOrder(orderinfo.AppID, orderinfo.Body, orderinfo.MchID, orderinfo.NonceStr, orderinfo.SpbillCreateIP,
		int(orderinfo.TotalFee), orderinfo.OutTradeNo, orderinfo.NotifyUrl, orderinfo.TradeType, orderinfo.OpenID, orderinfo.Key)
}

// PayCallback handle the response from wechat, call f(Out_trade_no,Transaction_id,Result_code) if sign is right
func PayCallback(ctx *server.Context, f func(string, string, string)) (string, string) {
	return wechat.WxpayCallback(ctx.Response(), ctx.Request(), f, conf.MMConf.AppOrderKey)
}

// HandleRecharge used as the second
func HandleRecharge(outTradeNo, transactionID, resultCode string) {
	outNum, err := strconv.Atoi(outTradeNo)
	if err != nil {
		log.Error("convert out_trade_no to number: " + err.Error())
	}
	switch resultCode {
	case "SUCCESS":
		err = model.RechargeService.Success(uint32(outNum), transactionID)
		if err != nil {
			log.Error("HandleRecharge: ", err)
		}
	case "FAIL":
		err = model.RechargeService.Fail(uint32(outNum))
		if err != nil {
			log.Error("HandleRecharge: ", err)
		}
	default:
		log.Error("unknown resultCode: " + resultCode)
	}
}

func CalculateSign(params map[string]interface{}, key string) string {
	sorted_keys := make([]string, 5)
	for k, _ := range params {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", params[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	return upperSign
}
