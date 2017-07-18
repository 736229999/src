package util

import "encoding/json"

//返回给前台的格式
type Response struct {
	Code int	 `json:"code"`	//返回码，具体见restful
	Data interface{} `json:"data"`	//返回的数据
}

func GetResponse(code int,data interface{}) string {
	response := new(Response)
	defer func() {
		response = nil
	}()
	response.Code = code
	response.Data = data
	content,err := json.Marshal(response)
	CheckError(err)
	return string(content)
}

func GetResponseMap(code int,data interface{})interface{}{
	returnMap := make(map[string]interface{})
	defer func() {
		returnMap = nil
	}()
	returnMap["code"] = code
	returnMap["data"] = data
	return returnMap
}