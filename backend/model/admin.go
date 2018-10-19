/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
 */

package model

import (
	"errors"

	"github.com/morgances/matchmaking/backend/conf"
)

type (
	adminServPrvd struct{}
)

var (
	AdminService adminServPrvd
)

func (adminServPrvd) Login(acc, pass string) error {
	row := DB.QueryRow(`SELECT COUNT(0) FROM `+conf.MMConf.Database+`.admin WHERE account=? AND password=? LOCK IN SHARE MODE`, acc, pass)
	var exist int64
	err := row.Scan(&exist)
	if err != nil || exist != 1 {
		return errors.New("login failed")
	}
	return nil
}
