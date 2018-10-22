/*
 * Revision History:
 *     Initial: 2018/10/22        Zhang Hao
 */

package util

import (
	"strings"
	"encoding/base64"
	"errors"
)

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
