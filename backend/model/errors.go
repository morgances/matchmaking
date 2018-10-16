/*
 * @Author: zhanghao
 * @DateTime: 2018-10-08 11:30:01
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:35:45
 */

package model

import (
	"errors"
)

var (
	ErrDuplicateEntry = errors.New("error duplicate entry")
	ErrMysql          = errors.New("error mysql")
)
