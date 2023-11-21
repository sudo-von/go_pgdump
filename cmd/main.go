package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sudo-von/go-pgdump/database"
)

var localDatabase = database.Database{
	Port:     os.Getenv("LOCAL_PORT"),
	Host:     os.Getenv("LOCAL_HOST"),
	User:     os.Getenv("LOCAL_USER"),
	Name:     os.Getenv("LOCAL_NAME"),
	Password: os.Getenv("LOCAL_PASSWORD"),
}

var remoteDatabase = database.Database{
	Port:     os.Getenv("REMOTE_PORT"),
	Host:     os.Getenv("REMOTE_HOST"),
	User:     os.Getenv("REMOTE_USER"),
	Name:     os.Getenv("REMOTE_NAME"),
	Password: os.Getenv("REMOTE_PASSWORD"),
}

func main() {

	schemaFilename := localDatabase.GetSchemaFilename()
	processSchema(schemaFilename)

	schema, err := localDatabase.GetSchema()
	if err != nil {
		log.Fatalln("Failed to get schema for", remoteDatabase.Name, ":", err.Error()+".")
	}

	tables := database.GetTables(schema)
	for _, table := range tables {

		trimmedTable := strings.Split(table, ".")
		schemaName := trimmedTable[0]
		tableName := trimmedTable[1]

		tableFilename := localDatabase.CreateTableFilename(schemaName, tableName)

		processTable(tableFilename, schemaName, tableName)
	}

	os.Exit(0)
}

func processSchema(filename string) {

	if err := localDatabase.CheckIfDatabaseExist(); err != nil {
		log.Fatalln("Local database", localDatabase.Name, "has not been created yet:", err.Error()+".")
	}

	if err := database.CheckIfFileExist(filename); err != nil {
		if err := remoteDatabase.DumpSchema(filename); err != nil {
			log.Fatalln("Failed to dump schema for", remoteDatabase.Name, ":", err.Error()+".")
		}
	}

	content := "PostgreSQL database dump complete"
	contains, err := database.CheckIfFileContains(filename, content)
	if err != nil {
		log.Fatalln("Failed to check if file", filename, "contains", content, ":", err.Error()+".")
	}

	if !contains {
		fmt.Println("File", filename, "does not contains", content+".")
		if err := remoteDatabase.DumpSchema(filename); err != nil {
			log.Fatalln("Failed to dump schema for", remoteDatabase.Name, ":", err.Error()+".")
		}
	}
}

func processTable(filename, schema, tablename string) {

	if err := database.CheckIfFileExist(filename); err != nil {
		if err := remoteDatabase.DumpTable(filename, schema, tablename); err != nil {
			log.Fatalln("Failed to dump table", tablename, "for", remoteDatabase.Name, ":", err.Error()+".")
		}
	}

	content := "PostgreSQL database dump complete"
	contains, err := database.CheckIfFileContains(filename, content)
	if err != nil {
		log.Fatalln("Failed to check if file", filename, "contains", content, ":", err.Error()+".")
	}

	if !contains {
		fmt.Println("File", filename, "does not contains", content+".")
		if err := remoteDatabase.DumpTable(filename, schema, tablename); err != nil {
			log.Fatalln("Failed to dump table", tablename, "for", remoteDatabase.Name, ":", err.Error()+".")
		}
	}
}
