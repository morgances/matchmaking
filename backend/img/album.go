/*
 * @Author: zhanghao
 * @Date: 2019-01-20 00:36:38
 * @Last Modified by: zhanghao
 * @Last Modified time: 2019-01-20 01:02:24
 */
package img

import (
	"net/http"
	"os"
)

func GetAlbum(id string) (imgs []string, err error) {
	dir := AlbumDir + id + "/"
	imgs, err = GetImages(dir)
	imgs = GetImageBase(imgs)
	for i := range imgs {
		imgs[i] = AlbumURL + imgs[i]
	}
	return
}

// SavePhotos save photos to user album
func SavePhotos(oid string, r *http.Request) error {
	return saveImages(AlbumDir+oid+"/", r)
}

func RemovePhotosIfExist(openid string, photos []string) {
	bases := GetImageBase(photos)
	if len(bases) == 0 {
		return
	}
	dir := AlbumDir + openid + "/"
	for _, base := range bases {
		os.Remove(dir + base)
	}
}
