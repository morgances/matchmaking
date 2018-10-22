/*
 * @Author: zhanghao
 * @DateTime: 2018-10-09 21:34:00
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 22:50:51
 */

package handler

import (
	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"github.com/morgances/matchmaking/backend/wx"
	"github.com/zh1014/comment/response"

	"github.com/silenceper/wechat/oauth"
	"github.com/dgrijalva/jwt-go"
)

type (
	userInfo struct {
		OpenID           string `json:"open_id"`
		NickName         string `json:"nick_name"`
		Sex              uint8  `json:"sex"`
		Age              uint8  `json:"age"`
		Height           string `json:"height"`
		Location         string `json:"location"`
		Job              string `json:"job"`
		Certified        bool   `json:"certified"`
		Vip              bool   `json:"vip"`
		SelfIntroduction string `json:"self_introduction"`
	}

	detailUserInfo struct {
		OpenID           string  `json:"open_id"`
		NickName         string  `json:"nick_name"`
		RealName         string  `json:"real_name"`
		Sex              uint8   `json:"sex"`
		Age              uint8   `json:"age"`
		Height           string  `json:"height"`
		Location         string  `json:"location"`
		Job              string  `json:"job"`
		Faith            string  `json:"faith"`
		Constellation    string  `json:"constellation"`
		SelfIntroduction string  `json:"self_introduction"`
		SelecCriteria    string  `json:"selec_criteria"`
		Certified        bool    `json:"certified"`
		Vip              bool    `json:"vip"`
		Points           float64 `json:"points"`
		Rose             uint32  `json:"rose"`
		Charm            uint32  `json:"charm"`
		DatePrivilege    uint32  `json:"date_privilege"`
	}

	fillInfo struct {
		Phone            string `json:"phone" validate:"required,numeric,len=11"`
		Wechat           string `json:"wechat" validate:"required"`
		RealName         string `json:"real_name" validate:"required"`
		Birthday         string `json:"birthday" validate:"required,len=10,contains=-"`
		Height           string `json:"height" validate:"required"`
		Job              string `json:"job" validate:"required"`
		Faith            string `json:"faith" validate:"required"`
		SelfIntroduction string `json:"self_introduction" validate:"required"`
		SelecCriteria    string `json:"selec_criteria" validate:"required"`
	}

	changeInfo struct {
		NickName         string `json:"nick_name" validate:"required"`
		Faith            string `json:"faith" validate:"required"`
		SelfIntroduction string `json:"self_introduction" validate:"required"`
		SelecCriteria    string `json:"selec_criteria" validate:"required"`
		Phone            string `json:"phone" validate:"required,numeric,len=11"`
		Wechat           string `json:"wechat" validate:"required"`
	}

	wechatCode struct {
		Code string `json:"code" validate:"required"`
	}

	targetOpenID struct {
		TargetOpenID string `json:"target_open_id" validate:"required,len=28"`
	}
)

var auth = wx.NewOauth()

// WechatLogin
func WechatLogin(this *server.Context) error {
	var (
		wechatCode wechatCode
		wechatData oauth.ResAccessToken
		userData   oauth.UserInfo
		resp       token
	)

	var err error
	if err = this.JSONBody(&wechatCode); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = this.Validate(&wechatCode); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	wechatData, err = auth.GetUserAccessToken(wechatCode.Code)
	if err != nil {
		log.Error("Error get user accessToken:", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatAuth, nil)
	}

	userData, err = auth.GetUserInfo(wechatData.AccessToken, wechatData.OpenID)
	if err != nil {
		log.Error("Error get user information:", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrWechatAuth, nil)
	}

	err = model.UserService.WeChatLogin(userData.OpenID, userData.Nickname, userData.City, uint8(userData.Sex))
	if err != nil {
		log.Error("wechat login failed: ", err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	util.SaveWechatAvatar(userData.OpenID, userData.HeadImgURL)

	resp.Token, err = util.NewToken(wechatData.OpenID, uint8(userData.Sex), false)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func FillInfo(this *server.Context) error {
	var (
		err    error
		req    fillInfo
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
	userp, err := model.UserService.FindByOpenID(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}

	age, err := util.GetAge(userp.Birthday)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	userp.Phone = req.Phone
	userp.Wechat = req.Wechat
	userp.Birthday = req.Birthday
	userp.Height = req.Height
	userp.Job = req.Job
	userp.RealName = req.RealName
	userp.Faith = req.Faith
	userp.SelfIntroduction = req.SelfIntroduction
	userp.SelecCriteria = req.SelecCriteria
	userp.Age = age

	if err = model.UserService.Update(userp); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func UserChangeInfo(this *server.Context) error {
	var (
		err error
		req changeInfo
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

	userp, err := model.UserService.FindByOpenID(openid)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}

	userp.NickName = req.NickName
	userp.Faith = req.Faith
	userp.SelfIntroduction = req.SelfIntroduction
	userp.SelecCriteria = req.SelecCriteria
	userp.Phone = req.Phone
	userp.Wechat = req.Wechat

	if err = model.UserService.Update(userp); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetUserDetail(this *server.Context) error {
	var (
		err   error
		req   targetOpenID
		userp *model.User
		resp  detailUserInfo
	)
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	userp, err = model.UserService.FindByOpenID(req.TargetOpenID)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}

	resp.OpenID = userp.OpenID
	resp.RealName = userp.RealName
	resp.Job = userp.Job
	resp.Height = userp.Height
	resp.SelecCriteria = userp.SelecCriteria
	resp.SelfIntroduction = userp.SelfIntroduction
	resp.Age = userp.Age
	resp.NickName = userp.NickName
	resp.Sex = userp.Sex
	resp.Charm = userp.Charm
	resp.Certified = userp.Certified
	resp.Points = userp.Points
	resp.Rose = userp.Rose
	resp.Vip = userp.Vip
	resp.Constellation = userp.Constellation
	resp.Location = userp.Location
	resp.Faith = userp.Faith
	resp.DatePrivilege = userp.DatePrivilege

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetRecommendUsers(this *server.Context) error {
	var (
		resp []userInfo
	)
	sex, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(uint8)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}

	// recommend contrast sex
	var allowedsex uint8 = 0
	if sex == 0 {
		allowedsex = 1
	} else {
		allowedsex = 0
	}

	userSlice, err := model.UserService.RecommendByCharm(allowedsex)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for i, user := range userSlice {
		resp = append(resp, userInfo{})
		resp[i].OpenID = user.OpenID
		resp[i].NickName = user.NickName
		resp[i].Sex = user.Sex
		resp[i].Age = user.Age
		resp[i].Height = user.Height
		resp[i].Location = user.Location
		resp[i].Job = user.Job
		resp[i].Certified = user.Certified
		resp[i].Vip = user.Vip
		resp[i].SelfIntroduction = user.SelfIntroduction
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetAlbum(this *server.Context) error {
	var (
		err          error
		isAbleToLook bool
		req          targetOpenID
		resp []string
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
	if openid != req.TargetOpenID {
		isAbleToLook, err = model.FollowService.FollowExist(openid, req.TargetOpenID)
		if err != nil {
			log.Error(err)
			return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
		}
	} else {
		isAbleToLook = true
	}

	if !isAbleToLook {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}

	resp, err = util.GetImages("./album/" + req.TargetOpenID + "/")
	if err != nil {
		if err == util.ErrNoImageExist {
			return response.WriteStatusAndDataJSON(this, constant.ErrNoAlbum, nil)
		}
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func UploadPhotos(this *server.Context) error {
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if err := util.SavePhotos(openid, this.Request()); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func RemovePhotos(this *server.Context) error {
	var (
		err error
		req struct {
			Images []string `json:"images" validate:"required,dive,required,contains=/"`
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
	util.RemovePhotosIfExist(openid, req.Images)
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func ChangeAvatar(this *server.Context) error {
	openid, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["open_id"].(string)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	req, _, err := this.Request().FormFile("avatar")
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	err = util.ChangeAvatar(openid, req)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func SendRose(this *server.Context) error {
	var (
		err error
		req struct {
			Reciever string `json:"reciever" validate:"required,len=28"`
			RoseNum  uint32 `json:"rose_num" validate:"required,gte=1"`
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

	if err = model.UserService.SendRose(openid, req.Reciever, req.RoseNum); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
