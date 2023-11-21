package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CheckIfFileExist(filename string) error {
	fmt.Println("Checking the existence of file", filename+".")

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Println("File", filename, "does not exist.")
		return err
	}

	fmt.Println("The file", filename, "was found.")
	return nil
}

func CheckIfFileContains(filename, content string) (bool, error) {
	fmt.Println("Checking if file", filename, "contains", content+".")

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	fileString := string(fileBytes)

	return strings.Contains(fileString, content), nil
}
