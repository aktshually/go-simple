package gosimple

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/End313234/go-simple/internal/utils/converters"
)

type Connection struct {
	Path     string // The path to the files
	Config   Config // Database config
	database Database
}

// Creates an instance of Connection
func (connection *Connection) New(path string, config Config) *Connection {
	connection.Path = path
	connection.database.Path = path

	connection.Config = config

	return connection
}

// Handles the connection to the database
func (connection *Connection) Connect(schemas ...Schema) error {
	dir, err := os.Stat(connection.Path)
	if err != nil {
		return errors.New("no permission to access the provided path")
	}

	if !dir.IsDir() {
		return errors.New("the provided path must be a valid directory")
	}

	flags := os.O_RDWR
	if connection.Config.CreateIfDoesNotExist {
		flags = os.O_CREATE | os.O_RDWR
	}

	for _, schema := range schemas {
		schemaName := reflect.TypeOf(&schema).Name()

		switch connection.Config.Pattern {
		case "PascalCase":
			break
		case "camelCase":
			schemaName = converters.ConvertPascalCaseToCamelCase(schemaName)
			break
		case "kebab-case":
			schemaName = converters.ConvertPascalCaseToKebabCase(schemaName)
			break
		case "snake_case":
			schemaName = converters.ConvertPascalCaseToSnakeCase(schemaName)
			break
		default:
			return errors.New("Pattern must be one of: PascalCase, camelCase, kebab-case, snake_case")
		}

		filePath := fmt.Sprintf("%s/%s.json", connection.Path, schemaName)

		file, err := os.OpenFile(filePath, flags, 0644)
		if err != nil {
			return errors.New("the file could not be created")
		}

		_, err = fmt.Fprint(file, "[]")
		if err != nil {
			return errors.New("could not write to the file")
		}

		connection.database.Schemas[&schema] = filePath
	}

	return nil
}

// Returns the database path and the schemas
func (connection Connection) GetDatabase() Database {
	return Database{
		Path:    connection.Path,
		Schemas: connection.database.Schemas,
	}
}
