/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package handler

import (
	"mime/multipart"
	"strconv"

	"net/http"

	"github.com/TechCatsLab/apix/http/server"
	log "github.com/TechCatsLab/logging/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/constant"
	"github.com/morgances/matchmaking/backend/img"
	"github.com/morgances/matchmaking/backend/model"
	"github.com/zh1014/comment/response"
)

type (
	goodsInfo struct {
		ID          uint32  `json:"id"`
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}
)

func CreateGoods(this *server.Context) error {
	// req form-data : price title(unique) description goods_image
	var (
		err    error
		lastId uint32
		resp   struct {
			GoodsID uint32 `json:"goods_id"`
		}
	)

	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}

	goods := &model.Goods{}
	goods.Price, err = strconv.ParseFloat(this.FormValue("price"), 64)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	goods.Title = this.FormValue("title")
	goods.Description = this.FormValue("description")

	lastId, err = model.GoodsService.Insert(goods)
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	resp.GoodsID = lastId

	image, _, err := this.Request().FormFile("goods_image")
	if err != nil {
		if err == http.ErrMissingFile {
			return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
		}
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, resp)
	}
	if err = img.SaveGoodsImage(lastId, image); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, resp)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func GetGoodsByID(this *server.Context) error {
	var (
		err   error
		req   targetID
		goods *model.Goods
		resp  goodsInfo
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
		resp   []goodsInfo
	)

	goodss, err = model.GoodsService.FindByPrice()
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	for _, goods := range goodss {
		var r goodsInfo
		r.ID = goods.ID
		r.Title = goods.Title
		r.Price = goods.Price
		r.Description = goods.Description
		resp = append(resp, r)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, resp)
}

func UpdateGoods(this *server.Context) error {
	var (
		req struct {
			ID          uint32  `json:"id" validate:"required,gte=1"`
			Title       string  `json:"title" validate:"required"`
			Price       float64 `json:"price" validate:"required,gte=1"`
			Description string  `json:"description"`
		}
		goods model.Goods
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	err := this.JSONBody(&req)
	if err != nil {
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
	goods.Description = req.Description
	if err = model.GoodsService.Update(&goods); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrMysql, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func ChangeGoodsImage(this *server.Context) error {
	var (
		err error
		gid int
		// req form-data: goods_id goods_image
		req struct {
			goodsID    uint32
			goodsImage multipart.File
		}
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	req.goodsImage, _, err = this.Request().FormFile("goods_image")
	if err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrInvalidParam, nil)
	}
	gid, err = strconv.Atoi(this.FormValue("goods_id"))
	req.goodsID = uint32(gid)

	if err = img.ChangeGoodsImage(req.goodsID, req.goodsImage); err != nil {
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSaveImage, nil)
	}

	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}

func DeleteGoods(this *server.Context) error {
	var (
		req targetID
	)
	isAdmin, ok := this.Request().Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["is_admin"].(bool)
	if !ok {
		return response.WriteStatusAndDataJSON(this, constant.ErrInternalServerError, nil)
	}
	if !isAdmin {
		return response.WriteStatusAndDataJSON(this, constant.ErrPermission, nil)
	}
	err := this.JSONBody(&req)
	if err != nil {
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
	if err = img.RemoveGoodsImage(req.TargetID); err != nil {
		// make a log but tell admin delete succeed, because it succeed in database
		log.Error(err)
		return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
	}
	return response.WriteStatusAndDataJSON(this, constant.ErrSucceed, nil)
}
