/*
 * @Author: zhanghao
 * @Date: 2019-01-20 00:58:23
 * @Last Modified by: zhanghao
 * @Last Modified time: 2019-01-20 01:03:05
 */
package img

import (
	"mime/multipart"
	"os"
	"strconv"
)

func ChangeGoodsImage(goodsid uint32, avatar multipart.File) error {
	return saveImage(GoodsDir+strconv.Itoa(int(goodsid)), avatar)
}

func RemoveGoodsImage(goodsid uint32) error {
	return os.RemoveAll(GoodsDir + strconv.Itoa(int(goodsid)))
}
