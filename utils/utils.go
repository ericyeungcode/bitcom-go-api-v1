package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func AToInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Fail to convert long: %v", s)
	}

	return val, nil
}

func MustAToInt64(s string) int64 {
	i, err := AToInt64(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

// bool field is always not empty
func IsEmptyQueryField(i interface{}) bool {
	switch v := i.(type) {
	case int:
		return v == 0
	case int64:
		return v == 0
	case float64:
		return v == 0.0
	case string:
		return v == ""
	default:
		return false
	}
}

func CombineQueryString(paramMap map[string]interface{}) string {
	var elementList []string
	for k, v := range paramMap {
		if !IsEmptyQueryField(v) {
			elementList = append(elementList, fmt.Sprintf("%v=%v", k, v))
		}
	}

	return strings.Join(elementList, "&")
}

func AnyToJsonStr(v any) string {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
