package converters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPascalCaseToKebabCase(t *testing.T) {
	assert := assert.New(t)

	s := "FooBar"
	s = ConvertPascalCaseToKebabCase(s)

	assert.Equal("foo-bar", s)
}
