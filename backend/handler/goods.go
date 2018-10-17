/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"fmt"
	"github.com/TechCatsLab/apix/http/server"
	"github.com/TechCatsLab/comment/response"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/morgances/matchmaking/backend/util"
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
			Price       int64  `json:"price" validate:"required,gte=0"`
			Description string `json:"description"`
		}
		lastId int64
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	image, _, err := this.Request().FormFile("goods_image")

	goods := &model.Goods{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
	}
	lastId, err = model.GoodsService.Insert(goods)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	if err = util.SaveImage(fmt.Sprintf("./goods/%d.jpg", lastId), image); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func GetGoodsByID(this *server.Context) error {
	var (
		err   error
		req   targetID
		goods *model.Goods
		resp  goodsResp
	)
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	goods, err = model.GoodsService.FindByID(req.TargetID)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	resp.ID = goods.ID
	resp.Title = goods.Title
	resp.Price = goods.Price
	resp.Description = goods.Description

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
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
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, goods := range goodss {
		var r goodsResp
		r.ID = goods.ID
		r.Title = goods.Title
		r.Price = goods.Price
		r.Description = goods.Description
		resp.Goods = append(resp.Goods, r)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func UpdateGoods(this *server.Context) error {
	var (
		err     error
		isAdmin bool
		req     struct {
			ID          int64  `json:"id" validate:"required,gte=1"`
			Title       string `json:"title" validate:"required"`
			Price       int64  `json:"price" validate:"required,gte=1"`
			Description string `json:"description"`
		}
		goods model.Goods
	)
	authorization := this.GetHeader("Authorization")
	_, _, _, isAdmin, err = util.ParseToken(authorization)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	goods.ID = req.ID
	goods.Title = req.Title
	goods.Price = req.Price
	goods.Description = goods.Description
	if err = model.GoodsService.Update(&goods); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
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
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	if err = this.JSONBody(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	if err = this.Validate(&req); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}

	if err = model.GoodsService.DeleteByID(req.TargetID); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
