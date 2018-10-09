/*
 * @Author: zhanghao
 * @Date: 2018-10-08 11:30:01
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-09 15:35:45
 */

package model

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("Err not found")
	ErrDuplicateEntry = errors.New("Err duplicate entry")
	ErrUnfollowFailed = errors.New("Err unfollow failed")
	ErrUserNotExist   = errors.New("Err user not exist")
)
