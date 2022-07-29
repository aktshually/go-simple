package converters

import (
	"strings"
)

func ConvertPascalCaseToCamelCase(s string) string {
	patternWithLowerSizedFirstCharacter := strings.ToLower(string(s[0]))
	return patternWithLowerSizedFirstCharacter + s[1:]
}
