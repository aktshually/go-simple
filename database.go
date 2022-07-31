package gosimple

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// Contains the database schemas and the path to the directory that
// stores the data
type Database struct {
	Path    string            // Path to the database
	Schemas map[Schema]string // Array containing the schemas
}

// Creates an instance of the schema in the database and returns it
func (database *Database) Create(schema Schema) (SchemaFile, error) {
	emptySchema := reflect.Zero(reflect.TypeOf(schema)).Interface()
	parsedFileContent := SchemaFile{}

	schemaFilePath, ok := database.Schemas[emptySchema]
	if !ok {
		return SchemaFile{}, errors.New("schema must be initialized")
	}

	content, err := ioutil.ReadFile(schemaFilePath)
	if err != nil {
		return SchemaFile{}, errors.New("could not read the schema file")
	}

	err = json.Unmarshal(content, &parsedFileContent)
	if err != nil {
		return SchemaFile{}, errors.New("invalid file content")
	}

	parsedFileContent = append(parsedFileContent, schema)

	file, err := os.OpenFile(schemaFilePath, os.O_RDWR, 0644)
	if err != nil {
		return SchemaFile{}, errors.New("could not open the file")
	}

	parsedSchema, _ := json.Marshal(parsedFileContent)
	fmt.Println(parsedSchema)
	fmt.Fprint(file, string(parsedSchema))

	return parsedFileContent, nil
}
