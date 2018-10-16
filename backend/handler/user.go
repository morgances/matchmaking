/*
 * @Author: zhanghao
 * @DateTime: 2018-10-09 21:34:00
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 22:50:51
 */

package handler

import (
	"log"
	"time"

	"github.com/silenceper/wechat/oauth"

	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"mime/multipart"
	"net/http"
	"strconv"
)

type (
	userInfo struct {
		OpenID           string `json: "open_id"`
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
		OpenID           string `json:"open_id"`
		NickName         string `json:"nick_name"`
		RealName         string `json:"real_name"`
		Sex              uint8  `json:"sex"`
		Age              uint8  `json:"age"` // todo: turn birthday to age before feedback
		Height           string `json:"height"`
		Location         string `json:"location"`
		Job              string `json:"job"`
		Faith            string `json:"faith"`
		Constellation    string `json:"constellation"`
		SelfIntroduction string `json:"self_introduction"`
		SelecCriteria    string `json:"selec_criteria"`
		Certified        bool   `json:"certified"`
		Vip              bool   `json:"vip"`
		Points           int64  `json:"points"`
		Rose             int64  `json:"rose"`
		Charm            int64  `json:"charm"`
	}

	fillInfo struct {
		Phone            string `json:"phone" validate:"required, numeric, len=11"`
		Wechat           string `json:"wechat" validate:"required"`
		RealName         string `json:"real_name" validate:"required"`
		Birthday         string `json:"birthday" validate:"required,len=10,contains=-"` // todo: validate xxxx-xx-xx format
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
		Phone            string `json:"phone" validate:"required, numeric, len=11"`
		Wechat           string `json:"wechat" validate:"required"`
	}

	wechatCode struct {
		Code string `json:"code" validate:"required"`
	}

	recharge struct {
		OpenID string `json:"open_id" validate:"required"`
		Rose   int64  `json:"rose" validate:"required"`
	}

	targetOpenID struct {
		TargetOpenID string `json:"target_open_id" validate:"required, len=28"`
	}
)

var auth = util.NewOauth()

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
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	if err = this.Validate(&wechatCode); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	// todo: need redirect ?
	if err = auth.Redirect(this.Response(), this.Request(), "127.0.0.1:3000/matchmaking/user/fillinfo", "", "301"); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrWechatAuth)
	}
	wechatData, err = auth.GetUserAccessToken(wechatCode.Code)
	if err != nil {
		log.Println("Error get user accessToken:", err)
		return this.WriteHeader(constant.ErrWechatAuth)
	}

	userData, err = auth.GetUserInfo(wechatData.AccessToken, wechatData.OpenID)
	if err != nil {
		log.Println("Error get user information:", err)
		return this.WriteHeader(constant.ErrWechatAuth)
	}

	err = model.UserService.WeChatLogin(userData.OpenID, userData.Nickname, userData.HeadImgURL, userData.City, uint8(userData.Sex))
	if err != nil {
		log.Println("Wechat login failed.")
		return this.WriteHeader(constant.ErrMysql)
	}

	resp.Token, err = util.NewToken(wechatData.OpenID, wechatData.AccessToken, uint8(userData.Sex), false)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrNewToken)
	}

	err = this.ServeJSON(resp)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(ErrNewToken)
	}
	return this.WriteHeader(http.StatusOK)
}

func FillInfo(this *server.Context) error {
	var (
		err      error
		req      fillInfo
		oid      string
		bornYear int
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

	userp, err := model.UserService.FindByOpenID(oid)
	if err != nil {
		return err
	}
	bornYear, err = strconv.Atoi(userp.Birthday[:3])
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	age := uint8(time.Now().Year() - bornYear)
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
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func UserChangeInfo(this *server.Context) error {
	var (
		err error
		req changeInfo
		oid string
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

	userp, err := model.UserService.FindByOpenID(oid)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	userp.NickName = req.NickName
	userp.Faith = req.Faith
	userp.SelfIntroduction = req.SelfIntroduction
	userp.SelecCriteria = req.SelecCriteria
	userp.Phone = req.Phone
	userp.Wechat = req.Wechat

	if err = model.UserService.Update(userp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetUserDetail(this *server.Context) error {
	var (
		err error
		req targetOpenID
		userp *model.User
		resp  detailUserInfo
	)
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	userp, err = model.UserService.FindByOpenID(req.TargetOpenID)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
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

	if err = this.ServeJSON(&resp); err != nil {
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetRecommendUsers(this *server.Context) error {
	var (
		err       error
		sex       uint8
		userSlice []model.User
		resp      struct {
			UserInformation []userInfo `json:"user_information"`
		}
	)

	authorization := this.GetHeader("Authorization")
	_, _, sex, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	// recommend contrast sex
	if sex == 0 {
		sex = 1
	} else {
		sex = 0
	}

	userSlice, err = model.UserService.RecommendByCharm(sex)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	for i, user := range userSlice {
		resp.UserInformation[i].OpenID = user.OpenID
		resp.UserInformation[i].NickName = user.NickName
		resp.UserInformation[i].Sex = user.Sex
		resp.UserInformation[i].Age = user.Age
		resp.UserInformation[i].Height = user.Height
		resp.UserInformation[i].Location = user.Location
		resp.UserInformation[i].Job = user.Job
		resp.UserInformation[i].Certified = user.Certified
		resp.UserInformation[i].Vip = user.Vip
		resp.UserInformation[i].SelfIntroduction = user.SelfIntroduction
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetAlbum(this *server.Context) error {
	var (
		err          error
		oid          string
		isAbleToLook bool
		req	targetOpenID
		resp struct {
			album []string `json:"album"`
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
	if oid != req.TargetOpenID {
		isAbleToLook, err = model.FollowService.FollowExist(oid, req.TargetOpenID)
		if err != nil {
			log.Println(err)
			return this.WriteHeader(constant.ErrMysql)
		}
	} else {
		isAbleToLook = true
	}

	if !isAbleToLook {
		return this.WriteHeader(constant.ErrNeedFollow)
	}

	resp.album, err = util.GetImages("./album/" + oid + "/")
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrLoadImage)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}

func UploadPhotos(this *server.Context) error {
	var (
		err error
		oid string
		req struct {
			ImageNum int `json:"image_num" validate:"required, numeric, gte=1"`
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
	if err = util.SavePhotos(req.ImageNum, oid, this.Request()); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrSaveImage)
	}
	return this.WriteHeader(http.StatusOK)
}

func RemovePhotos(this *server.Context) error {
	var (
		err error
		oid string
		req struct {
			Images []string `json:"images" validate:"required,dive,required"`
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
	util.RemovePhotosIfExist(oid, util.GetImageBase(req.Images))
	return this.WriteHeader(http.StatusOK)
}

func ChangeAvatar(this *server.Context) error {
	var (
		err error
		oid string
		req multipart.File
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	req, _, err = this.Request().FormFile("req")
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	err = util.ChangeAvatar(oid, req)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrServer)
	}
	return this.WriteHeader(http.StatusOK)
}

func SendRose(this *server.Context) error {
	var (
		req struct{
			Reciever string `json:"reciever" validate:"required,len=28"`
			RoseNum int `json:"rose_num" validate:"required,numeric,gte=1"`
		}
	)
	authorization := this.GetHeader("Authorization")
	oid, _, _, _, err := util.ParseToken(authorization)
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

	if err = model.UserService.SendRose(oid, req.Reciever, req.RoseNum); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

//func Recharge(this *server.Context) error {
//	var (
//		err error
//		recharge recharge
//	)
//	authorization := this.GetHeader("Authorization")
//	_, _, _, err = util.ParseToken(authorization)
//	if err != nil {
//		log.Println(err)
//		this.WriteHeader(constant.ErrBadJWT)
//	}
//	if err = this.JSONBody(&recharge); err != nil {
//		log.Println(err)
//		this.WriteHeader(constant.ErrInvalidParam)
//	}
//	if err = this.Validate(&recharge); err != nil {
//		log.Println(err)
//		this.WriteHeader(constant.ErrInvalidParam)
//	}
//
//	if err = model.UserService.Recharge(recharge.TargetOpenID, recharge.Rose); err != nil {
//		log.Println(err)
//		this.WriteHeader(constant.ErrMysql)
//	}
//	return this.WriteHeader(http.StatusOK)
//}

//func BecomeVIP(this *server.Context) error {
//	var (
//		err   error
//		oid   string
//	)
//	authorization := this.GetHeader("Authorization")
//	oid, _,_, err = util.ParseToken(authorization)
//	if err != nil {
//		log.Println(err)
//		return this.WriteHeader(constant.ErrInvalidParam)
//	}
//
//	if err = model.UserService.BecomeVIP(oid); err != nil {
//		log.Println(err)
//		return this.WriteHeader(constant.ErrInvalidParam)
//	}
//	return this.WriteHeader(http.StatusOK)
//}
