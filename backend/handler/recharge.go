/*
 * Revision History:
 *     Initial: 2018/10/19        Zhang Hao
 */

package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/wx"
	"github.com/zh1014/comment/response"
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
)

func RechargeVip(this *server.Context) error {
	var (
		resp struct {
			PrepayID string `json:"prepay_id"`
		}
	)
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	spbillCreateIp := this.Request().Header.Get(http.CanonicalHeaderKey("X-Forwarded-For"))

	outTradeNo, err := model.RechargeService.Insert("vip", openid, 1)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	unifyOrderResp, err := wx.SetOrder(wx.VipOrderInfo(spbillCreateIp, strconv.Itoa(int(outTradeNo)), openid))
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
	resp.PrepayID = unifyOrderResp.Prepay_id
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func RechargeRose(this *server.Context) error {
	var (
		err        error
		outTradeNo uint32
		req        struct {
			RoseNum uint32 `json:"rose_num" validate:"required,gte=1"`
		}
		resp struct {
			PrepayID string `json:"prepay_id"`
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

	spbillCreateIp := this.Request().Header.Get(http.CanonicalHeaderKey("X-Forwarded-For"))

	outTradeNo, err = model.RechargeService.Insert("rose", openid, req.RoseNum)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	unifyOrderResp, err := wx.SetOrder(wx.RoseOrder(spbillCreateIp, strconv.Itoa(int(outTradeNo)), openid, req.RoseNum))
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
	resp.PrepayID = unifyOrderResp.Prepay_id
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
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
