package converters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPascalCaseToSnakeCase(t *testing.T) {
	assert := assert.New(t)

	s := "FooBar"
	s = ConvertPascalCaseToSnakeCase(s)

	assert.Equal("foo_bar", s)
}
