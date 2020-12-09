package fileoperations

import (
	"io/ioutil"

	"log"

	"os"
)

//Create method to create file
func Create(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
}

//Read method to read file contents
func Read(fileName string) string {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	return string(fileData)
}

//Write method to write to a file
func Write(fileName string, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file : %s", err)
	}

	file.WriteString(data)
	defer file.Close()
}

//Append to append data at end of the file
func Append(fileName string, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Cannot open file : %s", err)
	}
	file.WriteString(data)
	defer file.Close()
}
