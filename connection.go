package gosimple

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/End313234/go-simple/internal/utils/converters"
)

// Creates an instance of Connection
func NewConnection(path string, config Config) *Connection {
	connection := Connection{
		Path:   path,
		Config: config,
		database: Database{
			Path:    path,
			Schemas: map[Schema]string{},
		},
	}

	return &connection
}

type Connection struct {
	Path     string // The path to the files
	Config   Config // Database config
	database Database
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

	var schemaName string
	for _, schema := range schemas {
		schemaName = reflect.TypeOf(schema).Name()
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
	}

	for _, schema := range schemas {
		schemaName := reflect.TypeOf(schema).Name()
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

		isFileExistent := false
		filePath := fmt.Sprintf("%s/%s.json", connection.Path, schemaName)
		_, err = os.Stat(filePath)
		if err == nil {
			isFileExistent = true
		}

		flags := os.O_RDWR
		if !isFileExistent {
			flags = os.O_CREATE | os.O_RDWR
		}

		file, err := os.OpenFile(filePath, flags, 0644)
		if err != nil {
			return errors.New("the file could not be created")
		}

		if !isFileExistent {
			_, err = fmt.Fprint(file, "[]")
			if err != nil {
				return errors.New("could not write to the file")
			}
		}

		connection.database.Schemas[schema] = filePath
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
