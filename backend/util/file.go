/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */

package util

import (
	"errors"
	"fmt"
	"github.com/morgances/matchmaking/backend/myerr"
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

func SaveImages(dir string, r *http.Request) error {
	num, err := strconv.Atoi(r.FormValue("image_num"))
	if err != nil {
		return errors.New("Save images: " + err.Error())
	}
	hasImageSaveFailed := false
	timeUnix := time.Now().Unix()
	for i := 1; i <= num; i++ {
		image, _, err := r.FormFile(fmt.Sprintf("image_%d", i))
		if err != nil {
			return errors.New(fmt.Sprintf("Save images %d: %v", i, err))
		}
		// todo: make images will not be created with the same name when one user upload photos twice in a second
		// todo: should I return err when one of images failed to save ?
		err = SaveImage(dir+fmt.Sprintf("%d-%d.jpg", timeUnix, i), image)
		if err != nil {
			hasImageSaveFailed = true
		}
		image.Close()
	}
	if hasImageSaveFailed {
		return errors.New("There is image failed to saved in "+dir)
	}
	return nil
}

func SavePhotos(oid string, r *http.Request) error {
	return SaveImages("./album/"+oid+"/", r)
}

func SavePostImages(id int, r *http.Request) error {
	return SaveImages(fmt.Sprintf("./post/%d/", id), r)
}

func ChangeAvatar(oid string, avatar multipart.File) error {
	return SaveImage("./avatar/"+oid+".jpg", avatar)
}

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

// RemoveImageIfExist todo: need update images of post ?
func RemoveImageIfExist(name string) error {
	err := os.Remove(name)
	if os.IsNotExist(err) || err == nil {
		return nil
	}
	return err
}

func RemovePhotosIfExist(oid string, bases []string) {
	dir := "./album/" + oid + "/"
	for _, base := range bases {
		os.Remove(dir + base)
	}
}

func GetImages(dir string) (imgs []string, err error) {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	var infos []os.FileInfo
	infos, err = ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, myerr.ErrUserHasNoAlbum
		}
		return nil, err
	}
	for _, info := range infos {
		imgs = append(imgs, dir+info.Name())
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
