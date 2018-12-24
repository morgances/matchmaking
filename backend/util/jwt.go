/*
 * Revision History:
 *     Initial: 2018/10/22        Zhang Hao
 */

package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/morgances/matchmaking/backend/conf"
)

func NewToken(oid string, sex uint8, isAdm bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"open_id":  oid,
		"sex":      sex,
		"is_admin": isAdm,
		"exp":      time.Now().Add(30 * 24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(conf.MMConf.PrivateTokenKey))
}
