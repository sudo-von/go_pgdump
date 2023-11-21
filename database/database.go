package database

import (
	"fmt"
	"os/exec"
)

type Database struct {
	Port     string
	Host     string
	User     string
	Name     string
	Password string
}

func (d Database) CheckIfDatabaseExist() error {
	fmt.Println("Checking if database", d.Name, "exist.")

	command := fmt.Sprintf(`psql "%s" -c "\q"`, d.GetConnectionString())

	cmd := exec.Command("bash", "-c", command)

	_, err := cmd.CombinedOutput()
	return err
}

func (d Database) GetConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", d.User, d.Password, d.Host, d.Port, d.Name)
}

func (d Database) GetSchema() ([]byte, error) {
	fmt.Println("Getting schema for", d.Name+".")

	cmd := exec.Command("psql", d.GetConnectionString(), "-c", "SELECT CONCAT(table_schema,'.',table_name) FROM information_schema.tables WHERE table_schema IN('public', 'global') AND table_name NOT IN ('pg_stat_statements_info', 'pg_stat_statements');", "-A", "-t", "-q")
	stdout, err := cmd.CombinedOutput()
	return stdout, err
}

func (d Database) GetSchemaFilename() string {
	return fmt.Sprintf("./files/%s_schema.sql", d.Name)
}

func (d Database) CreateTableFilename(schema, tablename string) string {
	return fmt.Sprintf("./files/%s/%s_table.sql", schema, tablename)
}
