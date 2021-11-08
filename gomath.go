package gomath

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// MaxFloat 取最大值,返回float64类型
//
// 参数: a,b两个数值类型数据
func MaxFloat(a, b interface{}) (float64, error) {
	var aFloat64, bFloat64 float64

	aValue := reflect.ValueOf(a)
	if aValue.CanConvert(reflect.TypeOf(float64(aFloat64))) {
		aFloat64 = aValue.Convert(reflect.TypeOf(aFloat64)).Float()
	} else {
		return 0, errors.New("类型错误")
	}

	bValue := reflect.ValueOf(b)
	if bValue.CanConvert(reflect.TypeOf(bFloat64)) {
		bFloat64 = bValue.Convert(reflect.TypeOf(bFloat64)).Float()
	} else {
		return 0, errors.New("类型错误")
	}

	if aFloat64 < bFloat64 {
		return bFloat64, nil
	} else {
		return aFloat64, nil
	}
}

// MaxInt 取最大值,返回int64类型
//
// 参数: a,b两个数值类型数据
func MaxInt(a, b interface{}) (int64, error) {
	var aInt, bInt int64
	aValue := reflect.ValueOf(a)
	if aValue.CanConvert(reflect.TypeOf(aInt)) {
		aInt = aValue.Convert(reflect.TypeOf(aInt)).Int()
	} else {
		return 0, errors.New("a: 类型错误")
	}

	bValue := reflect.ValueOf(b)
	if bValue.CanConvert(reflect.TypeOf(bInt)) {
		bInt = bValue.Convert(reflect.TypeOf(bInt)).Int()
	} else {
		return 0, errors.New("b: 类型错误")
	}

	if aInt > bInt {
		return aInt, nil
	} else {
		return bInt, nil
	}
}

// MinFloat 取最小值,返回float64类型
//
// 参数: a,b两个数值类型数据
func MinFloat(a, b interface{}) (float64, error) {
	var aFloat64, bFloat64 float64

	aValue := reflect.ValueOf(a)
	if aValue.CanConvert(reflect.TypeOf(float64(aFloat64))) {
		aFloat64 = aValue.Convert(reflect.TypeOf(aFloat64)).Float()
	} else {
		return 0, errors.New("类型错误")
	}

	bValue := reflect.ValueOf(b)
	if bValue.CanConvert(reflect.TypeOf(bFloat64)) {
		bFloat64 = bValue.Convert(reflect.TypeOf(bFloat64)).Float()
	} else {
		return 0, errors.New("类型错误")
	}

	if aFloat64 < bFloat64 {
		return aFloat64, nil
	} else {
		return bFloat64, nil
	}
}

// MinInt 取最大值,返回int64类型
//
// 参数: a,b两个数值类型数据
func MinInt(a, b interface{}) (int64, error) {
	var aInt, bInt int64
	aValue := reflect.ValueOf(a)
	if aValue.CanConvert(reflect.TypeOf(aInt)) {
		aInt = aValue.Convert(reflect.TypeOf(aInt)).Int()
	} else {
		return 0, errors.New("a: 类型错误")
	}

	bValue := reflect.ValueOf(b)
	if bValue.CanConvert(reflect.TypeOf(bInt)) {
		bInt = bValue.Convert(reflect.TypeOf(bInt)).Int()
	} else {
		return 0, errors.New("b: 类型错误")
	}

	if aInt > bInt {
		return bInt, nil
	} else {
		return aInt, nil
	}
}

// RoundFloat (四舍五入)保留多少位小数
//
// 只支持float32, float64
func RoundFloat(source interface{}, decimal int) (dest float64, err error) {
	sourceValue := reflect.ValueOf(source)
	var destTemp float64
	if sourceValue.CanConvert(reflect.TypeOf(float64(dest))) {
		destTemp = sourceValue.Convert(reflect.TypeOf(float64(dest))).Float()
	} else {
		return 0, errors.New("非数值类型不能执行该操作")
	}
	dest, err = strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", decimal), destTemp), 64)
	return
}

// ArrayToInt 数组按位转换为int
func ArrayToInt(source []int) (result int, err error) {
	for _, v := range source {
		if v > 31 || v == 0 {
			return 0, errors.New("类型越界")
		}
		result += 1 << v
	}
	return
}

// IntToArray int按位转换为数组
func IntToArray(source int) (result []int, err error) {
	for index := 1; index < 31; index++ {
		if (source & 1 << index) == 1 {
			result = append(result, index)
		}
	}
	return
}

// ArrayToStr 数组按位转换为int
func ArrayToStr(source []int) (result string, err error) {
	b, err := json.Marshal(source)
	return string(b), err
}

// StrToArray 字符串转数组
func StrToArray(source string) (result []int, err error) {
	err = json.Unmarshal([]byte(source), &result)
	return
}
