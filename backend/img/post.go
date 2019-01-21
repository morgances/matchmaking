/*
 * @Author: zhanghao
 * @Date: 2019-01-20 00:57:26
 * @Last Modified by: zhanghao
 * @Last Modified time: 2019-01-20 01:02:30
 */
package img

import (
	"net/http"
	"os"
	"strconv"
)

func SavePostImages(id uint32, r *http.Request) error {
	return saveImages(PostDir+strconv.Itoa(int(id))+"/", r)
}

// ClearPostImages if exist
func ClearPostImages(postid uint32) error {
	err := os.RemoveAll(PostDir + strconv.Itoa(int(postid)) + "/")
	if err != nil {
		return err
	}
	return nil
}

func GetPostImgs(id uint32) []string {
	dir := PostDir + strconv.Itoa(int(id)) + "/"
	imgs, _ := GetImages(dir)
	if imgs == nil {
		return nil
	}
	for i := range imgs {
		imgs[i] = PostURL + imgs[i]
	}
	return imgs
}
