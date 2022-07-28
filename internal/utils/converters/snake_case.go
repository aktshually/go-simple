package converters

import (
	"fmt"
	"strings"
)

func ConvertPascalCaseToSnakeCase(s string) string {
	s = ConvertPascalCaseToCamelCase(s)

	for _, r := range s {
		char := string(r)
		lowerSizedChar := strings.ToLower(char)
		if lowerSizedChar != char {
			s = strings.ReplaceAll(s, char, fmt.Sprintf("_%s", lowerSizedChar))
		}
	}

	return s
}
