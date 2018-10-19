/*
 * @Author: zhanghao
 * @DateTime: 2018-10-09 22:16:30
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-10 02:03:29
 */

package wx

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/morgances/matchmaking/backend/conf"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/oauth"
)

var (
	auth *oauth.Oauth

	errParseToken = errors.New("err parse token")
)

const (
	TokenKeyOpenID   = "open_id"
	TokenKeyAccessID = "access_id"
	TokenKeySex      = "sex"
	TokenKeyIsAdmin  = "is_admin"
)

func NewToken(oid, acid string, sex uint8, isAdm bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"open_id":   oid,
		"access_id": acid,
		"sex":       sex,
		"is_admin":  isAdm,
		"exp":       time.Now().Add(30 * 24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(conf.MMConf.PrivateTokenKey))
}

func ParseToken(tokenString string) (oid, acid string, sex uint8, isAdm bool, err error) {
	var sexFloat64 float64

	kv := strings.Split(tokenString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		err = errors.New("invalid token authorization string")
		return "", "", 0, false, err
	}
	tokenString = kv[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errParseToken
		}

		return []byte(conf.MMConf.PrivateTokenKey), nil
	})
	if err != nil {
		return "", "", 0, false, err
	}
	tokenMap, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if oid, ok = tokenMap[TokenKeyOpenID].(string); !ok {
			return "", "", 0, false, errParseToken
		}
		if acid, ok = tokenMap[TokenKeyAccessID].(string); !ok {
			return "", "", 0, false, errParseToken
		}
		if sexFloat64, ok = (tokenMap[TokenKeySex].(float64)); !ok {
			fmt.Println(reflect.TypeOf(tokenMap[TokenKeySex]), ":", tokenMap[TokenKeySex])
			return "", "", 0, false, errParseToken
		}
		if isAdm, ok = tokenMap[TokenKeyIsAdmin].(bool); !ok {
			return "", "", 0, false, errParseToken
		}
		return oid, acid, uint8(sexFloat64), isAdm, nil
	}
	return "", "", 0, false, errParseToken
}

func NewOauth() *oauth.Oauth {
	return oauth.NewOauth(&context.Context{
		AppID:     conf.MMConf.AppID,
		AppSecret: conf.MMConf.AppID,
	})
}

func ParseBase64(auth string) (acc, pass string, err error) {
	if !strings.HasPrefix(auth, "Basic ") {
		return "", "", errors.New("ParseBase64: wrong authorization way")
	}
	auth = strings.TrimPrefix(auth, "Basic ")

	authBytes, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return "", "", err
	}
	authSlice := strings.SplitN(string(authBytes), ":", 2)

	if len(authSlice) != 2 {
		return "", "", errors.New("error lack password: " + auth)
	}
	return authSlice[0], authSlice[1], nil
}
