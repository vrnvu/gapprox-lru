package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

func main() {
	fileSize := int64(1e3)
	maxNumberOfFiles := int64(10)
	limitSize := fileSize * maxNumberOfFiles
	fmt.Println(limitSize)

	workingPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	directory := path.Join(workingPath, "/tmp")
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		fmt.Println(fmt.Sprintf("{name: %s, size: %d}", f.Name(), f.Size()))
	}

	err = clean(directory, limitSize)
	if err != nil {
		log.Fatal(err)
	}

}

func clean(directory string, limitSize int64) error {
	rand.Seed(time.Now().UnixNano())

	files := getFiles(directory)
	currentSize := getCurrentSize(files)

	for limitSize < currentSize {
		switch {
		case len(files) == 0:
			return errors.New("cannot free enough disk space")
		case len(files) == 1:
			os.Remove(path.Join(directory, files[0].Name()))
			currentSize -= files[0].Size()
			fmt.Printf("removed %s\n", files[0].Name())
			files = nil
		default:
			index_a, index_b := getRandomIndexes(len(files))
			a := files[index_a]
			b := files[index_b]

			if a.ModTime().UnixNano() < b.ModTime().UnixNano() {
				os.Remove(path.Join(directory, b.Name()))
				currentSize -= b.Size()
				files = append(files[:index_b], files[index_b+1:]...)
				fmt.Printf("removed %s\n", b.Name())
			} else {
				os.Remove(path.Join(directory, a.Name()))
				currentSize -= a.Size()
				files = append(files[:index_a], files[index_a+1:]...)
				fmt.Printf("removed %s\n", a.Name())
			}
		}
	}

	return nil
}

// getRandomIndexes returns two indexes among the files to perform the algorithm
// this number could be parametrize, for intance this is a paramter of redis
func getRandomIndexes(numberOfFiles int) (int, int) {
	index_a := rand.Intn(numberOfFiles)
	index_b := rand.Intn(numberOfFiles)

	// in case of collision get another index
	for index_a == index_b {
		index_b = rand.Intn(numberOfFiles)
	}

	return index_a, index_b

}

// getFiles assumes no nested folders inside directory
func getFiles(directory string) []os.FileInfo {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err)
	}

	return files

}

// getCurrentSize returns the size, assuming no recursive folders for this example
func getCurrentSize(files []os.FileInfo) int64 {
	var total int64
	for _, f := range files {
		total += f.Size()
	}
	return total

}
