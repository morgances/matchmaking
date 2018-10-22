/*
 * @Author: zhanghao
 * @DateTime: 2018-10-10 01:52:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 02:00:46
 */

package router

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/handler"
)

var (
	Router *server.Router
	Skip   = [...]string{
		"/matchmaking/user/wechatlogin",
		"/matchmaking/admin/login",
		constant.NotifyUrl,
	}
)

func init() {
	Router = server.NewRouter()

	Router.Post("/matchmaking/user/wechatlogin", handler.WechatLogin)
	Router.Post("/matchmaking/user/fillinfo", handler.FillInfo)
	Router.Post("/matchmaking/user/modifyinfo", handler.UserChangeInfo)
	Router.Post("/matchmaking/user/changeavatar", handler.ChangeAvatar)
	Router.Post("/matchmaking/user/uploadphotos", handler.UploadPhotos)
	Router.Post("/matchmaking/user/removephotos", handler.RemovePhotos)
	Router.Post("/matchmaking/user/sendrose", handler.SendRose)
	Router.Post("/matchmaking/user/album", handler.GetAlbum)           									// todo: use Get
	Router.Post("/matchmaking/user/userdetail", handler.GetUserDetail) 		 // both user and admin		 	todo: use Get
	Router.Get("/matchmaking/user/recommendusers", handler.GetRecommendUsers)

	Router.Post("/matchmaking/comment/insert", handler.CommentService.Insert)
	Router.Post("/matchmaking/comment/change", handler.CommentService.ChangeContent)
	Router.Post("/matchmaking/comment/ofuser", handler.CommentService.ListCommentsByUserID) 				// todo: use Get
	Router.Post("/matchmaking/comment/ofpost", handler.CommentService.ListCommentsByTarget) 				// todo: use Get

	Router.Post("/matchmaking/follow/follow", handler.Follow)
	Router.Post("/matchmaking/follow/unfollow", handler.Unfollow)
	Router.Get("/matchmaking/follow/following", handler.GetFollowing)
	Router.Get("/matchmaking/follow/follower", handler.GetFollower)

	Router.Post("/matchmaking/goods/byid", handler.GetGoodsByID)      		 // both user and admin 		todo: use Get
	Router.Get("/matchmaking/goods/byprice", handler.GetGoodsByPrice) 		 // both user and admin

	Router.Post("/matchmaking/post/create", handler.CreatePost)
	Router.Post("/matchmaking/post/commend", handler.CommendPost)
	Router.Post("/matchmaking/post/delete?isadmin=false", handler.DeletePost)
	Router.Get("/matchmaking/post/many?isreviewed=true", handler.GetReviewedPost)  		// both user and admin
	Router.Get("/matchmaking/post/mine", handler.GetMyPost)

	Router.Post("/matchmaking/signin/signin", handler.Signin)
	Router.Get("/matchmaking/signin/mysigninrecord", handler.GetSigninRecord)

	Router.Post("/matchmaking/trade/create", handler.CreateTrade)
	Router.Get("/matchmaking/trade/mytrades", handler.GetMyTrades)

	Router.Post("/matchmaking/recharge/vip", handler.RechargeVip)
	Router.Post("/matchmaking/recharge/rose", handler.RechargeRose)

	// for admin below ====================================================================
	Router.Post("/matchmaking/admin/login", handler.Login)
	Router.Post("/matchmaking/user/certifypass", handler.Certify)
	Router.Post("/matchmaking/user/dateprivilegereduce", handler.DatePrivilegeReduce)
	Router.Post("/matchmaking/user/dateprivilegeadd", handler.DatePrivilegeAdd)
	Router.Post("/matchmaking/user/contacts", handler.GetContact)											 // todo: use Get ?

	Router.Get("/matchmaking/trade/unfinished", handler.GetUnfinishedTrade)
	Router.Post("/matchmaking/trade/cancel", handler.CancelTrade)
	Router.Post("/matchmaking/trade/updatestatus", handler.UpdateTradeStatus)

	Router.Get("/matchmaking/post/many?isreviewed=false", handler.GetUnreviewedPost)
	Router.Post("/matchmaking/post/updatestatus", handler.UpdatePostStatus)
	Router.Post("/matchmaking/post/delete?isadmin=true", handler.AdminDeletePost)

	Router.Post("/matchmaking/goods/create", handler.CreateGoods)
	Router.Post("/matchmaking/goods/update", handler.UpdateGoods)
	Router.Post("/matchmaking/goods/changeimage", handler.ChangeGoodsImage)
	Router.Post("/matchmaking/goods/delete", handler.DeleteGoods)

	Router.Get("/matchmaking/recharge/record", handler.GetRechargeRecord)

	// for wechat notify ====================================================================
	Router.Post(constant.NotifyUrl, handler.PayResult)
}
