package mapUtils

import (
	"strings"
)

func GetCaseInsensitiveKey(m map[string]string, searchKey string) (string, bool) {
	lowerSearchKey := strings.ToLower(searchKey)
	for key, value := range m {
		if strings.ToLower(key) == lowerSearchKey {
			return value, true
		}
	}
	return "", false
}
