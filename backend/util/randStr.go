/*
 * Revision History:
 *     Initial: 2018/12/23        Zhang Hao
 */

package util

import (
	"math/rand"
	"time"
)

const stdChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomStr(strLen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte,strLen)
	for i := 0; i<strLen; i++ {
		bytes[i] = stdChars[r.Intn(len(stdChars))]
	}
	return string(bytes)
}