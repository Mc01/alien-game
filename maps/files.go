package main

import (
	"bufio"
	"os"
)

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
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

func openWriter(file *os.File) *bufio.Writer {
	return bufio.NewWriter(file)
}

func writeLine(writer *bufio.Writer, line string) {
	_, err := writer.WriteString(line + "\n")
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
