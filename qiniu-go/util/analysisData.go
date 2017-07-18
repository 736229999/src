package util

import (
	"encoding/json"
)
//数据解析工具类

/**
	解析json数据
 */
func AnalysisJson(jsonData string)map[string]interface{}  {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonData),&jsonMap)
	CheckError(err)
	return jsonMap
}
