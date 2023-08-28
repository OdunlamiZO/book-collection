package util

import (
	"fmt"
	"strings"
)

func ArrayToString(array []string) string {
	quotedArray := make([]string, len(array))
	for index, item := range array {
		quotedArray[index] = fmt.Sprintf(`"%s"`, item)
	}
	quotedString := fmt.Sprintf("{%s}", strings.Join(quotedArray, ", "))
	return quotedString
}

func StringToArray(str string) []string {
	quotedArray := strings.Split(strings.Trim(str, "{}"), ", ")
	array := make([]string, len(quotedArray))
	for index, item := range quotedArray {
		array[index] = strings.Trim(item, `"`)
	}
	return array
}
