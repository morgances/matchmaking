/*
 * Revision History:
 *     Initial: 2018/10/16        Zhang Hao
 */

package util

import (
	"strconv"
	"strings"
	"errors"
)

// GetConstellation get the constellation of date, but it do not know the date that not exist
func GetConstellation(date string) (string, error) {
	dateSlice := strings.SplitN(date, "-", 3)
	month, err := strconv.Atoi(dateSlice[1])
	if err != nil {
		return "", errors.New("invalid date")
	}
	day, err := strconv.Atoi(dateSlice[2])
	if err != nil {
		return "", errors.New("invalid date")
	}
	index := month
	constellationArry := [13]string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座","双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	sepDay := []int{20,19,21,21,21,22,23,23,23,23,22,22}
	if day < sepDay[index] {
		return constellationArry[index-1], nil
	}
	return constellationArry[index], nil
}
