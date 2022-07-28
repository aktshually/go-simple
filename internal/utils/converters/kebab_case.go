package converters

import (
	"fmt"
	"strings"
)

func ConvertPascalCaseToKebabCase(s string) string {
	s = ConvertPascalCaseToCamelCase(s)

	for _, r := range s {
		char := string(r)
		lowerSizedChar := strings.ToLower(char)
		if lowerSizedChar != char {
			s = strings.ReplaceAll(s, char, fmt.Sprintf("-%s", lowerSizedChar))
		}
	}

	return s
}
