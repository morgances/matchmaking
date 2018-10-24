/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var ErrNoImageExist = errors.New("error user has no album")

// SaveWechatAvatar save the wechat avatar as the initial avatar
func SaveWechatAvatar(oid, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create("./avatar/" + oid + ".jpg")
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir("./avatar", 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func ChangeAvatar(oid string, avatar multipart.File) error {
	return SaveImage("./avatar/"+oid+".jpg", avatar)
}

// SavePhotos save photos to user album
func SavePhotos(oid string, r *http.Request) error {
	return SaveImages("./album/"+oid+"/", r)
}

func RemovePhotosIfExist(openid string, photos []string) {
	bases := GetImageBase(photos)
	if len(bases) == 0 {
		return
	}
	dir := "./album/" + openid + "/"
	for _, base := range bases {
		os.Remove(dir + base)
	}
}

// operate post image -----------------------------------------------

func SavePostImages(id uint32, r *http.Request) error {
	return SaveImages(fmt.Sprintf("./post/%d/", id), r)
}

// ClearPostImages if exist
func ClearPostImages(postid uint32) error {
	err := os.RemoveAll(fmt.Sprintf("./post/%d", postid))
	if err != nil {
		return errors.New(fmt.Sprintf("ClearPostImages ./post/%d/ :", postid) + err.Error())
	}
	return nil
}

// operate goods image--------------------------------------------------

func ChangeGoodsImage(goodsid uint32, avatar multipart.File) error {
	return SaveImage(fmt.Sprintf("./goods/%d.jpg", goodsid), avatar)
}

func RemoveGoodsImage(goodsid uint32) error {
	return os.RemoveAll(fmt.Sprintf("./goods/%d.jpg", goodsid))
}

// ------------------------------------------------------------------

func SaveImages(dir string, r *http.Request) error {
	numString := r.FormValue("image_num")
	if numString == "" { // no image
		return nil
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		return errors.New("Save images: " + err.Error())
	}
	hasImageSaveFailed := true
	timeUnix := time.Now().Unix()
	for i := 1; i <= num; i++ {
		image, _, err := r.FormFile(fmt.Sprintf("image_%d", i))
		if err != nil {
			return errors.New(fmt.Sprintf("Save images %d: %v", i, err))
		}
		// todo: let images will not be created with the same name when one user upload photos twice in a second
		// todo: should I return err when one of images failed to save ?
		err = SaveImage(dir+fmt.Sprintf("%d-%d.jpg", timeUnix, i), image)
		if err != nil {
			hasImageSaveFailed = true
		}
		image.Close()
	}
	if hasImageSaveFailed {
		return errors.New("There is image failed to be saved in " + dir)
	}
	return nil
}

// SaveImage cover image if it already exist
func SaveImage(name string, image multipart.File) error {
	localImage, err := os.Create(name)
	defer localImage.Close()
	if err != nil {
		if os.IsNotExist(err) {
			lastSlash := strings.LastIndex(name, "/")
			if lastSlash < 0 {
				return errors.New("error directory with out slash")
			}
			if err = os.MkdirAll(name[:lastSlash], 0755); err != nil {
				return errors.New("Save image: error make directory " + name[:lastSlash] + " :" + err.Error())
			}

			// retry to create file
			if localImage, err = os.Create(name); err != nil {
				return errors.New("SaveImage: " + err.Error())
			}
		} else {
			return errors.New("Save image: error create file: " + name)
		}
	}

	if _, err = io.Copy(localImage, image); err != nil {
		return errors.New("Save image: error io copy: " + err.Error())
	}
	return nil
}

func GetImages(dir string) (imgs []string, err error) {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	var infos []os.FileInfo
	infos, err = ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoImageExist
		}
		return nil, err
	}
	for _, info := range infos {
		name := info.Name()
		if !strings.HasSuffix(name, ".jpg") {
			continue
		}
		imgs = append(imgs, dir+name)
	}
	if imgs == nil {
		return nil, ErrNoImageExist
	}
	return imgs, nil
}

// GetImageBase get bases of an arry of paths, path is skipped when it is empty or consists entirely of slashes
func GetImageBase(paths []string) (base []string) {
	for _, aPath := range paths {
		aPath = strings.TrimRight(aPath, "/ ")
		if len(aPath) == 0 {
			continue
		}
		base = append(base, path.Base(aPath))
	}
	return base
}
