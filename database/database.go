package database

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/End313234/go-simple/internal/utils/converters"
)

type Database struct {
	Path    string
	Schemas []any
}

/* Handles the connection to the database. If `createIfDoesNotExist`
*  is set to `true`, the file will be create if the provided file
*  does not exist. However, if set to `false`, the function will panic
*  if the file does not exist.
 */
func (database *Database) Connect(config *Config, schemas ...Schema) error {
	dir, err := os.Stat(database.Path)
	if err != nil {
		return errors.New("no permission to access the provided path")
	}

	if !dir.IsDir() {
		return errors.New("the provided path must be a valid directory")
	}

	flags := os.O_RDWR
	if config.CreateIfDoesNotExist {
		flags = os.O_CREATE | os.O_RDWR
	}

	for _, schema := range schemas {
		schemaName := reflect.TypeOf(schema).Name()

		switch config.Pattern {
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

		file, err := os.OpenFile(fmt.Sprintf("%s/%s.json", database.Path, schemaName), flags, 0644)
		if err != nil {
			return errors.New("the file could not be created")
		}

		_, err = fmt.Fprint(file, "[]")
		if err != nil {
			return errors.New("could not write to the file")
		}
	}

	return nil
}
