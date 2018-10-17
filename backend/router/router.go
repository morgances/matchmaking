/*
 * @Author: zhanghao
 * @DateTime: 2018-10-10 01:52:13
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 02:00:46
 */

package router

import (
	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/handler"
)

var (
	Router *server.Router
	Skip   = [...]string{
		"/matchmaking/user/wechatlogin",
		"/matchmaking/admin/login",
	}
)

func init() {
	Router = server.NewRouter()

	Router.Post("/matchmaking/user/wechatlogin", handler.WechatLogin)
	Router.Post("/matchmaking/user/fillinfo", handler.FillInfo)
	Router.Post("/matchmaking/user/userchangeinfo", handler.UserChangeInfo)
	Router.Post("/matchmaking/user/uploadphotos", handler.UploadPhotos)
	Router.Post("/matchmaking/user/removephotos", handler.RemovePhotos)
	Router.Post("/matchmaking/user/changeavatar", handler.ChangeAvatar)
	Router.Post("/matchmaking/user/sendrose", handler.SendRose)
	Router.Post("/matchmaking/user/getalbum", handler.GetAlbum)           // todo: Get
	Router.Post("/matchmaking/user/getuserdetail", handler.GetUserDetail) // todo: Get
	Router.Get("/matchmaking/user/getrecommendusers", handler.GetRecommendUsers)

	Router.Post("/matchmaking/comment/create", handler.CommentService.Insert)
	Router.Post("/matchmaking/comment/change", handler.CommentService.ChangeContent)
	Router.Get("/matchmaking/comment/commentsofuser", handler.CommentService.ListCommentsByUserID)
	Router.Get("/matchmaking/comment/commentsofpost", handler.CommentService.ListCommentsByTarget)

	Router.Post("/matchmaking/follow/follow", handler.Follow)
	Router.Post("/matchmaking/follow/unfollow", handler.Unfollow)
	Router.Get("/matchmaking/follow/getfollowing", handler.GetFollowing)
	Router.Get("/matchmaking/follow/getfollower", handler.GetFollower)

	Router.Post("/matchmaking/goods/getgoodsbyid", handler.GetGoodsByID) // both user and admin todo: Get
	Router.Get("/matchmaking/goods/getgoodsbyprice", handler.GetGoodsByPrice)

	Router.Post("/matchmaking/post/createpost", handler.CreatePost) // todo:test util here 10.18
	Router.Post("/matchmaking/post/commendpost", handler.CommendPost)
	Router.Post("/matchmaking/post/deletepost", handler.DeletePost)
	Router.Get("/matchmaking/post/getreviewedpost", handler.GetReviewedPost)
	Router.Get("/matchmaking/post/getmypost", handler.GetMyPost)

	Router.Post("/matchmaking/signin/signin", handler.Signin)
	Router.Get("/matchmaking/post/getsignrecord", handler.GetSigninRecord)

	Router.Post("/matchmaking/post/createtrade", handler.CreateTrade)
	Router.Get("/matchmaking/post/getmytrades", handler.GetMyTrades)

	// for admin below ----------------------------------------------------------------------
	Router.Post("/matchmaking/admin/login", handler.Login)
	Router.Post("/matchmaking/admin/certify", handler.Certify)
	Router.Post("/matchmaking/admin/dateprivilegereduce", handler.DatePrivilegeReduce)
	Router.Post("/matchmaking/admin/dateprivilegeadd", handler.DatePrivilegeAdd)
	Router.Post("/matchmaking/admin/canceltrade", handler.CancelTrade)
	Router.Post("/matchmaking/admin/updatetradestatus", handler.UpdateTradeStatus)
	Router.Post("/matchmaking/admin/updatepoststatus", handler.UpdatePostStatus)
	Router.Post("/matchmaking/admin/admindeletepost", handler.AdminDeletePost)
	Router.Post("/matchmaking/admin/getcontact", handler.GetContact) // todo: Get
	Router.Get("/matchmaking/admin/getunfinishedtrade", handler.GetUnfinishedTrade)
	Router.Get("/matchmaking/admin/getunreviewedpost", handler.GetUnreviewedPost)

	Router.Post("/matchmaking/goods/creategoods", handler.CreateGoods) // admin
	Router.Post("/matchmaking/goods/updategoods", handler.UpdateGoods) // admin
	Router.Post("/matchmaking/goods/deletegoods", handler.DeleteGoods) // admin
}
