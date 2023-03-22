package sutil

import (
	"encoding/json"
	"math/rand"
	"reflect"

	"github.com/json-iterator/go"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary

	defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

//RandomString(6)
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func NotEmy(s string) bool {
	return !(len(s) == 0)
}

func JsonString(v interface{}) string {
	Json, _ := Json.Marshal(v)
	return string(Json)
}

func JsonUnmarshalFromString(jsonStr string, v interface{}) error {
	err := Json.UnmarshalFromString(jsonStr, v)
	if err != nil {
		return json.Unmarshal([]byte(jsonStr), v)
	}
	return err
}

func InStringSlice(val string, ss []string) bool {
	if len(val) == 0 {
		return false
	}
	for _, s := range ss {
		if val == s {
			return true
		}
	}
	return false
}

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}