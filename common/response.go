package common

func Success[T any](data T) map[string]any {
	var resultMap = make(map[string]any, 3)
	resultMap["code"] = 0
	resultMap["message"] = "请求成功"
	resultMap["data"] = data

	return resultMap
}

func Fail[T any](message string) map[string]any {
	var resultMap = make(map[string]any, 3)
	resultMap["code"] = -1
	resultMap["message"] = message
	resultMap["data"] = nil

	return resultMap
}
