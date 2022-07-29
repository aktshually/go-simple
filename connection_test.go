package gosimple

import (
	"fmt"
	"os"
	"testing"

	"github.com/End313234/go-simple/constants"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.RemoveAll("./internal/databases")
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	connectionPath := "./internal/databases"
	connectionConfig := Config{
		CreateIfDoesNotExist: true,
		Pattern:              constants.KEBAB_CASE,
	}

	connection := NewConnection(connectionPath, connectionConfig)

	assert.Equal(connectionPath, connection.Path)
	assert.Equal(connectionConfig, connection.Config)
	assert.Equal(connectionPath, connection.database.Path)
	assert.Empty(connection.database.Schemas)
}

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	type UserSchema struct {
		Id   string
		Name string
	}

	futureUserSchemaFilePath := "./internal/databases/user-schema.json"
	connection := NewConnection("./internal/databases", Config{
		CreateIfDoesNotExist: true,
		Pattern:              constants.KEBAB_CASE,
	})
	err := connection.Connect(UserSchema{})

	assert.NoError(err)
	assert.FileExists(futureUserSchemaFilePath)
	assert.Equal(map[Schema]string{
		UserSchema{}: futureUserSchemaFilePath,
	}, connection.database.Schemas)
}

func TestGetDatabase(t *testing.T) {
	assert := assert.New(t)
	db := Database{}
	databasePath := "./internal/databases"
	futureUserSchemaFilePath := fmt.Sprintf("%s/user-schema.json", databasePath)

	type UserSchema struct {
		Id   string
		Name string
	}

	connection := NewConnection(databasePath, Config{
		CreateIfDoesNotExist: true,
		Pattern:              constants.KEBAB_CASE,
	})
	connection.Connect(UserSchema{})
	db = connection.GetDatabase()

	assert.Equal(databasePath, db.Path)
	assert.Equal(map[Schema]string{
		UserSchema{}: futureUserSchemaFilePath,
	}, db.Schemas)
}
