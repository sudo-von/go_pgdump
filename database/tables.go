package database

import (
	"strings"
)

func GetTables(schemaBytes []byte) []string {
	schemaString := string(schemaBytes)
	trimmedSchema := strings.TrimSpace(schemaString)
	tables := strings.Split(trimmedSchema, "\n")
	return tables
}
