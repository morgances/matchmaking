/*
 * @Author: zhanghao
 * @Date: 2019-01-20 00:58:23
 * @Last Modified by: zhanghao
 * @Last Modified time: 2019-01-20 01:03:05
 */
package img

import (
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
)

func ChangeGoodsImage(id uint32, f multipart.File) error {
	return SaveGoodsImage(id, f)
}

func RemoveGoodsImage(goodsid uint32) error {
	return os.RemoveAll(GoodsDir + strconv.Itoa(int(goodsid)))
}

func SaveGoodsImage(id uint32, f multipart.File) error {
	n := fmt.Sprintf(GoodsDir+"%d.jpg", id)
	return saveImage(n, f)
}
