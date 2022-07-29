package go_simple

// Contains the database schemas and the path to the directory that
// stores the data
type Database struct {
	Path    string             // Path to the database
	Schemas map[*Schema]string // Array containing the schemas
}
