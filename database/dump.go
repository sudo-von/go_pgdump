package database

import (
	"fmt"
	"os/exec"
)

func (d Database) DumpSchema(filename string) error {
	fmt.Println("Dumping schema for", d.Name, "database.")

	command := fmt.Sprintf("pg_dump %s --schema-only -f %s", d.GetConnectionString(), filename)

	cmd := exec.Command("bash", "-c", command)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println("Schema for", d.Name, "database was successfully dumped.")
	return nil
}

func (d Database) DumpTable(filename, schema, tablename string) error {
	fmt.Println("Dumping table", tablename, "for", d.Name, "database.")

	command := fmt.Sprintf(`pg_dump %s -t '%s."%s"' --data-only -f %s`, d.GetConnectionString(), schema, tablename, filename)

	cmd := exec.Command("bash", "-c", command)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println("Table for", d.Name, "database was successfully dumped.")
	return nil
}
