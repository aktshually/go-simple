package converters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPascalCaseToCamelCase(t *testing.T) {
	assert := assert.New(t)

	s := "FooBar"
	s = ConvertPascalCaseToCamelCase(s)

	assert.Equal("fooBar", s)
}
