package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Sign(key string, path string, paramMap map[string]any) string {
	if path == "" || paramMap == nil {
		return ""
	}

	hash := hmac.New(sha256.New, []byte(key))
	var sb strings.Builder
	sb.WriteString(path)
	sb.WriteString("&")
	sb.WriteString(encodeMap(paramMap))

	strToSign := sb.String()

	//hash.Reset()
	hash.Write([]byte(strToSign))
	sha := hex.EncodeToString(hash.Sum(nil))
	return sha
}

func encodeMapList(itemList []map[string]any) string {
	var strList []string
	for _, v := range itemList {
		objVal := encodeMap(v)
		objStr := fmt.Sprintf("%v", objVal)
		strList = append(strList, objStr)
	}

	outputStr := strings.Join(strList, "&")
	outputStr = "[" + outputStr + "]"
	return outputStr
}

func encodeMap(paramMap map[string]any) string {
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	resultMap := map[string]any{}

	for _, k := range keys {
		value := paramMap[k]
		switch v := value.(type) {
		case []map[string]any:
			listStr := encodeMapList(v)
			resultMap[k] = listStr
		case map[string]any:
			mapStr := encodeMap(v)
			resultMap[k] = mapStr
		case bool:
			boolStr := strconv.FormatBool(v)
			resultMap[k] = boolStr
		default:
			generalStr := fmt.Sprintf("%v", v)
			resultMap[k] = generalStr
		}
	}

	return buildParamToSign(resultMap)
}

func buildParamToSign(paramMap map[string]any) string {
	var keys []string
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultStrList []string

	for _, k := range keys {
		var sb strings.Builder
		sb.WriteString(k)
		sb.WriteString("=")
		strValue := fmt.Sprintf("%v", paramMap[k])
		sb.WriteString(strValue)
		resultStrList = append(resultStrList, sb.String())
	}

	outputStr := strings.Join(resultStrList, "&")
	return outputStr
}
