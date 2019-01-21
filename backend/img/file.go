/*
 * Revision History:
 *     Initial: 2018/10/15        Zhang Hao
 */
package img

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

	log "github.com/TechCatsLab/logging/logrus"
	"github.com/morgances/matchmaking/backend/conf"
)

var (
	AlbumDir  = conf.MMConf.ImgRoot + "album/"
	AvatarDir = conf.MMConf.ImgRoot + "avatar/"
	PostDir   = conf.MMConf.ImgRoot + "post/"
	GoodsDir  = conf.MMConf.ImgRoot + "goods/"

	AlbumURL  = conf.MMConf.ImgURLPrefix + "album/"
	AvatarURL = conf.MMConf.ImgURLPrefix + "avatar/"
	PostURL   = conf.MMConf.ImgURLPrefix + "post/"
	GoodsURL  = conf.MMConf.ImgURLPrefix + "goods/"
)

var ErrNoImageExist = errors.New("error user has no album")

// SaveWechatAvatar save the wechat avatar as the initial avatar
func SaveWechatAvatar(oid, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(AvatarDir + oid + ".jpg")
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(AvatarDir, 0755); err != nil {
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

func ChangeAvatar(oid string, f multipart.File) error {
	return saveImage(AvatarDir+oid+".jpg", f)
}

func saveImages(dir string, r *http.Request) error {
	numString := r.FormValue("image_num")
	// return nil when there is no image
	if numString == "" {
		return nil
	}

	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Error("saveImages: " + err.Error())
		return nil
	}
	for i := 1; i <= num; i++ {
		f, _, err := r.FormFile(fmt.Sprintf("image_%d", i))
		if err != nil {
			continue
		}
		na := dir + fmt.Sprintf("%d.jpg", time.Now().UnixNano())
		err = saveImage(na, f)
		if err != nil {
			log.Error("saveImages " + na + " failed: " + err.Error())
		}
		f.Close()
	}
	return nil
}

// SaveImage will cover origin image when name is the same
func saveImage(name string, image multipart.File) error {
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
	imgs = make([]string, 0, len(infos))
	for _, info := range infos {
		imgs = append(imgs, info.Name())
	}
	if len(imgs) == 0 {
		return nil, ErrNoImageExist
	}
	return imgs, nil
}

// GetImageBase get bases of an arry of paths
func GetImageBase(paths []string) (base []string) {
	if paths == nil {
		return nil
	}
	base = make([]string, 0, len(paths))
	for _, aPath := range paths {
		if aPath == "" {
			continue
		}
		base = append(base, path.Base(aPath))
	}
	return base
}
