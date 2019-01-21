/*
 * @Author: zhanghao
 * @Date: 2019-01-20 00:57:26
 * @Last Modified by: zhanghao
 * @Last Modified time: 2019-01-20 01:02:30
 */
package img

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func SavePostImages(id uint32, r *http.Request) error {
	return saveImages(PostDir+strconv.Itoa(int(id)), r)
}

// ClearPostImages if exist
func ClearPostImages(postid uint32) error {
	err := os.RemoveAll(PostDir + strconv.Itoa(int(postid)))
	if err != nil {
		return errors.New(fmt.Sprintf("ClearPostImages /post/%d/ :", postid) + err.Error())
	}
	return nil
}

func GetPostImgs(postid uint32) (imgs []string, err error) {
	dir := PostDir + strconv.Itoa(int(postid))
	imgs, err = GetImages(dir)
	imgs = GetImageBase(imgs)
	for i := range imgs {
		imgs[i] = AlbumURL + imgs[i]
	}
	return
}
