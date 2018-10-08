/*
 * @Author: zhanghao
 * @Date: 2018-10-08 11:30:01
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-10-08 11:30:48
 */

package model

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}

type DuplicateError struct {
	Err error
}

func (e DuplicateError) Error() string {
	return e.Err.Error()
}
