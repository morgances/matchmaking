/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"fmt"
	"github.com/TechCatsLab/apix/http/server"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
	"log"
	"net/http"
)

type (
	goodsResp struct {
		ID          int64  `json:"id"`
		Title       string `json:"title"`
		Price       int64  `json:"price"`
		Description string `json:"description"`
	}
)

func CreateGoods(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     struct {
			Title       string `json:"title" validate:"required"`
			Price       int64  `json:"price" validate:"required, numeric, gte=0"`
			Description string `json:"description"`
		}
		lastId int64
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	image, _, err := this.Request().FormFile("goods_image")

	goods := &model.Goods{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
	}
	lastId, err = model.GoodsService.Insert(goods)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	if err = util.SaveImage(fmt.Sprintf("./goods/%d.jpg", lastId), image); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrSaveImage)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetGoodsByID(this *server.Context) error {
	var (
		err   error
		req   targetID
		goods *model.Goods
		resp  goodsResp
	)
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	goods, err = model.GoodsService.FindByID(req.TargetID)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	resp.ID = goods.ID
	resp.Title = goods.Title
	resp.Price = goods.Price
	resp.Description = goods.Description
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}

func GetGoodsByPrice(this *server.Context) error {
	var (
		err    error
		goodss []model.Goods
		resp   struct {
			Goods []goodsResp `json:"goods"`
		}
	)

	goodss, err = model.GoodsService.FindByPrice()
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	for _, goods := range goodss {
		var r goodsResp
		r.ID = goods.ID
		r.Title = goods.Title
		r.Price = goods.Price
		r.Description = goods.Description
		resp.Goods = append(resp.Goods, r)
	}
	if err = this.ServeJSON(&resp); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	return this.WriteHeader(http.StatusOK)
}

func UpdateGoods(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     struct {
			ID          int64  `json:"id" validate:"required, numeric, gte=1"`
			Title       string `json:"title" validate:"required"`
			Price       int64  `json:"price" validate:"required, numeric, gte=1"`
			Description string `json:"description"`
		}
		goods model.Goods
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}

	goods.ID = req.ID
	goods.Title = req.Title
	goods.Price = req.Price
	goods.Description = goods.Description
	if err = model.GoodsService.Update(&goods); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}

func DeleteGoods(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     targetID
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrInvalidParam)
	}
	if !isAdmin {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}
	if err = this.Validate(&req); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrPermissionDenied)
	}

	if err = model.GoodsService.DeleteByID(req.TargetID); err != nil {
		log.Println(err)
		return this.WriteHeader(constant.ErrMysql)
	}
	return this.WriteHeader(http.StatusOK)
}
