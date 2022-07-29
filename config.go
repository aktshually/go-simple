package gosimple

type Config struct {
	CreateIfDoesNotExist bool   // If truthy, creates the file in the provided pattern if the file does not exist
	Pattern              string // The pattern for the file name. Must be one of: PascalCase, camelCase, kebab-case or snake_case
}
