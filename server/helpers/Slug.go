package helpers

import (
	"strings"
)

func CreateSlug(name string) string {
	newValue := strings.ToLower(name)
	newValue = strings.ReplaceAll(newValue, " ", "-")
	return newValue
}
