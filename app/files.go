package main

import "os"

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

