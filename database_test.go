package gosimple

import (
	"testing"

	"github.com/End313234/go-simple/constants"
	"github.com/stretchr/testify/assert"
)

func TestCreateDocumentSuccess(t *testing.T) {
	assert := assert.New(t)

	type UserSchema struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	connection := NewConnection("./internal/databases", Config{
		Pattern: constants.KEBAB_CASE,
	})
	connection.Connect(UserSchema{})

	db := connection.GetDatabase()
	doc, err := db.Create(UserSchema{
		Id:   1,
		Name: "foo",
	})

	assert.NoError(err)
	assert.NotEmpty(doc)
}

func TestCreateDocumentFailure(t *testing.T) {
	assert := assert.New(t)

	type UserSchema struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	type PersonSchema struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	connection := NewConnection("./internal/databases", Config{
		Pattern: constants.KEBAB_CASE,
	})
	connection.Connect(UserSchema{})

	db := connection.GetDatabase()
	doc, err := db.Create(PersonSchema{
		Id:   1,
		Name: "foo",
	})

	assert.Error(err)
	assert.Empty(doc)
}
