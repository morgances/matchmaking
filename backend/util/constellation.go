/*
 * Revision History:
 *     Initial: 2018/10/16        Zhang Hao
 */

package util

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// GetConstellation get the constellation of date, but it do not know the date that not exist
func GetConstellation(date string) (string, error) {
	dateSlice := strings.SplitN(date, "-", 3)
	month, err := strconv.Atoi(dateSlice[1])
	if err != nil {
		return "", errors.New("invalid date :" + date)
	}
	day, err := strconv.Atoi(dateSlice[2])
	if err != nil {
		return "", errors.New("invalid date :" + date)
	}
	index := month
	constellationArry := [13]string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	sepDay := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	if day < sepDay[index] {
		return constellationArry[index-1], nil
	}
	return constellationArry[index], nil
}

func GetAge(date string) (uint8, error) {
	if len(date) != 10 {
		return 0, errors.New("GetAge: invalid date :" + date)
	}
	year, err := strconv.Atoi(date[:4])
	if err != nil {
		return 0, errors.New("GetAge: " + err.Error())
	}

	return uint8(time.Now().Year() - year), nil
}

func GetAgeAndConstell(date string) (uint8, string, error) {
	age, err := GetAge(date)
	if err != nil {
		return 0, "", err
	}
	constell, err := GetConstellation(date)
	if err != nil {
		return 0, "", err
	}
	return age, constell, nil
}
