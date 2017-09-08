package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// GetRespOrError 用Get方式获取返回
func GetRespOrError(url string, result interface{}, validName string, validData interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%v", string(b))
		return
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return
	}
	validField := getValue(result, validName)
	if validData != nil && validField.Interface() != validData {
		err = fmt.Errorf("%v", string(b))
		return
	}
	if isEmptyValue(validField) {
		err = fmt.Errorf("%v", string(b))
	}
	return
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	if v.Type() == reflect.TypeOf(time.Time{}) {
		return v.Interface().(time.Time).IsZero()
	}

	return false
}

func getValue(field interface{}, validName string) reflect.Value {
	v := reflect.ValueOf(field)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			break
		}
		v = v.Elem()
	}
	validField := v.FieldByName(validName)
	if !validField.IsValid() {
		for i := 0; i < v.NumField(); i++ {
			tempField := v.Field(i)
			for tempField.Kind() == reflect.Ptr {
				if tempField.IsNil() {
					break
				}
				tempField = v.Field(i).Elem()
			}
			if tempField.Kind() == reflect.Struct {
				iLoopField := tempField.FieldByName(validName)
				if iLoopField.IsValid() {
					validField = iLoopField
				}
			}
		}
	}
	return validField
}
