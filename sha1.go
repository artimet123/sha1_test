package main

import (
    "os"
    "fmt"
    "crypto/sha1"
    "io"
)

func hashFile(path string) (hashName string,err error) {
	file, err := os.Open(path)
	h := sha1.New()

	if err != nil {
		return "",err
	}

	defer file.Close()

	_,err = io.Copy(h,file)
	if err != nil {
		return "",err
	}


	hashName = fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(path)
	return hashName,nil
}
