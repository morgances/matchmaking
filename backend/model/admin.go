/*
 * Revision History:
 *     Initial: 2018/10/13        Zhang Hao
 */

package model

type (
	adminServPrvd struct{}
)

var (
	AdminService adminServPrvd
)

func (adminServPrvd) Login(acc, pass string) error {
	row := DB.QueryRow(`SELECT * FROM admin WHERE account=?, password=? LOCK IN SHARE MODE`, acc, pass)
	var exist int64
	err := row.Scan(&exist)
	return err
}
