package env

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// 读取配置文件
func Get(key string, defaultValue interface{}) interface{} {

	// 读取 JSON 文件
	data, err := ioutil.ReadFile("env.json")
	if err != nil {
		return defaultValue
	}

	// 解析 JSON 内容到 map[string]interface{} 类型
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return defaultValue
	}

	// 使用.分割下级字段路径
	fields := strings.Split(key, ".")

	// 逐级访问下级字段
	for _, field := range fields {
		value, ok := result[field]
		if !ok {
			return defaultValue
		}

		result, ok = value.(map[string]interface{})
		if !ok {
			return value
		}
	}

	return result
}
