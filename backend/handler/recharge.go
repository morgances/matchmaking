/*
 * Revision History:
 *     Initial: 2018/10/19        Zhang Hao
 */

package handler

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/conf"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"github.com/morgances/matchmaking/backend/wx"
	"github.com/zh1014/comment/response"

	"github.com/193Eric/go-wechat"
)

const (
	// Re-signing is consistent with the signature type of the unified order
	// The third-part package I chosen using MD5 to sign for unify order
	signType = "MD5"
)

type (
	recharge struct {
		ID            uint32
		OpenID        string
		Project       string
		Num           uint32
		Fee           uint32
		TransactionID string
		Status        uint8
	}

	signAgainResp struct {
		AppID     string `json:"app_id"`
		TimeStamp string `json:"time_stamp"`
		NonceStr  string `json:"nonce_str"`
		Package   string `json:"package"`
		SignType  string `json:"sign_type"`
		PaySign   string `json:"pay_sign"`
	}
)

func RechargeVip(this *server.Context) error {
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	outTradeNo, err := model.RechargeService.Insert("vip", openid, 1)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	unifyOrderResp, err := wx.SetOrder(wx.VipOrderInfo(util.RemoteIp(this.Request()), strconv.Itoa(int(outTradeNo)), openid))
	if err != nil {
		log.Error(err)
		if err = model.RechargeService.Fail(outTradeNo); err != nil {
			log.Error(err)
		}
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatPay, nil)
	}
	if unifyOrderResp.Result_code != "SUCCESS" {
		log.Errorf("RechargeVip: recharge id=%d, return_code=%s, result_code=%s", outTradeNo, unifyOrderResp.Return_code, unifyOrderResp.Result_code)
		if err = model.RechargeService.Fail(outTradeNo); err != nil {
			log.Error(err)
		}
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatPay, nil)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, *signAgain(unifyOrderResp))
}

func RechargeRose(this *server.Context) error {
	var (
		err        error
		outTradeNo uint32
		req        struct {
			RoseNum uint32 `json:"rose_num" validate:"required,gte=1"`
		}
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

	outTradeNo, err = model.RechargeService.Insert("rose", openid, req.RoseNum)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	unifyOrderResp, err := wx.SetOrder(wx.RoseOrder(util.RemoteIp(this.Request()), "123456789_"+strconv.Itoa(int(outTradeNo)), openid, req.RoseNum))
	if err != nil {
		log.Error(err)
		if err = model.RechargeService.Fail(outTradeNo); err != nil {
			log.Error(err)
		}
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatPay, nil)
	}
	if unifyOrderResp.Result_code != "SUCCESS" {
		log.Errorf("RechargeRose: recharge id=%d, return_code=%s, result_code=%s", outTradeNo, unifyOrderResp.Return_code, unifyOrderResp.Result_code)
		if err = model.RechargeService.Fail(outTradeNo); err != nil {
			log.Error(err)
		}
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatPay, nil)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, *signAgain(unifyOrderResp))
}

func signAgain(uoResp *wechat.UnifyOrderResp) *signAgainResp {
	resp := new(signAgainResp)
	resp.AppID = uoResp.Appid
	resp.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
	resp.NonceStr = uoResp.Nonce_str
	resp.Package = "prepay_id=" + uoResp.Prepay_id
	resp.SignType = signType

	params := make(map[string]interface{}, 5)
	params["appId"] = resp.AppID
	params["timeStamp"] = resp.TimeStamp
	params["nonceStr"] = resp.NonceStr
	params["package"] = resp.Package
	params["signType"] = resp.SignType

	resp.PaySign = wx.CalculateSign(params, conf.MMConf.AppOrderKey)
	params = nil
	return resp
}

func GetRechargeRecord(this *server.Context) error {
	var (
		resp []recharge
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}

	rchgs, err := model.RechargeService.FindAll()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, rchg := range rchgs {
		oneRecord := recharge{}
		oneRecord.ID = rchg.ID
		oneRecord.OpenID = rchg.OpenID
		oneRecord.Project = rchg.Project
		oneRecord.Num = rchg.Num
		oneRecord.Fee = rchg.Fee
		oneRecord.TransactionID = rchg.TransactionID
		oneRecord.Status = rchg.Status
		resp = append(resp, oneRecord)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func PayResult(this *server.Context) error {
	outTradeNo, result := wx.PayCallback(this, wx.HandleRecharge)
	if result == "FAIL" {
		return errors.New("PayResult: wechat pay callback failed")
	}
	fmt.Println("recharge id:" + outTradeNo + " callback succeed")
	return nil
}
