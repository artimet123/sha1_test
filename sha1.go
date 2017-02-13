package main

import (
    "os"
    "fmt"
    "path/filepath"
    "crypto/sha1"
    "io"
)

func hashFile(path string) (hashName string,error) {
	file, err := os.Open(path)
	h := sha1.New()

	if err != nil {
		return err
	}

	defer file.Close()

	data := make([]byte, 1024)
	for{
		n, err := file.Read(data){
		if n != 0 {
			io.WriteString(h, string(data)) 
		} else {
			break
		}

		if err != nil && err != io.EOF {
			panic(err) //Å×³öÒì³£
		}
	}

	hashName := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(path)
	return nil 
}
