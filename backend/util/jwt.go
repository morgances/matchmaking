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

//func ParseToken(tokenString string) (oid string, sex uint8, isAdm bool, err error) {
//	var sexFloat64 float64
//
//	kv := strings.Split(tokenString, " ")
//	if len(kv) != 2 || kv[0] != "Bearer" {
//		err = errors.New("invalid token authorization string")
//		return "", 0, false, err
//	}
//	tokenString = kv[1]
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errParseToken
//		}
//
//		return []byte(conf.MMConf.PrivateTokenKey), nil
//	})
//	if err != nil {
//		return "", 0, false, err
//	}
//	tokenMap, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		if oid, ok = tokenMap[TokenKeyOpenID].(string); !ok {
//			return "", 0, false, errParseToken
//		}
//		if sexFloat64, ok = (tokenMap[TokenKeySex].(float64)); !ok {
//			fmt.Println(reflect.TypeOf(tokenMap[TokenKeySex]), ":", tokenMap[TokenKeySex])
//			return "", 0, false, errParseToken
//		}
//		if isAdm, ok = tokenMap[TokenKeyIsAdmin].(bool); !ok {
//			return "", 0, false, errParseToken
//		}
//		return oid, uint8(sexFloat64), isAdm, nil
//	}
//	return "", 0, false, errParseToken
//}
