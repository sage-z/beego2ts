package beego2ts

import (
	"fmt"
	"os"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func type2ts(t string) string {
	switch t {
	case "string":
		return "string"
	case "integer":
		return "number"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	default:
		fmt.Println("未知的类型 ", t)
		return "any"
	}
}
