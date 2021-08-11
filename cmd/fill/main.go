package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	workingPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	directory := path.Join(workingPath, "/tmp")

	numberOfFiles := 13

	_, err = os.Stat(directory)
	if os.IsNotExist(err) {
		os.Mkdir(directory, 0777)
		fmt.Printf("%s created\n", directory)
	}

	for i := 0; i < numberOfFiles; i++ {
		fileName := fmt.Sprintf("%d.txt", i)
		filePath := path.Join(directory, fileName)
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		if err := f.Truncate(1e3); err != nil {
			log.Fatal(err)
		}
	}
}
