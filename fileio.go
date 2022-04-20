package main

import (
	"fmt"
	"os"
)

func slurpInMem(filename string) error {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(string(dat))
	return nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	return file, err
}

// read only first n bytes
func readBytes(file os.File, n int) error {
	buf := make([]byte, n)
	n1, err := file.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes: %s\n", n1, string(buf[:n1]))
	return nil
}

func seekAndRead(file os.File, n int) error {
	_, err := file.Seek(6, 0)
	return err
}
