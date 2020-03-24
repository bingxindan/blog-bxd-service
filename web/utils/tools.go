package utils

import (
	"encoding/json"
	"log"
)

/**
 * json转换为map
 * @time 2020-03-24
 * @author zm
 */
func JsonToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		log.Panic(err)
	}

	return tempMap
}