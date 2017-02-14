package main

import (
"fmt"
"io/ioutil"
"os"
"strings"
)

func ListDir(dirPth string, suffix string)(err error) {
	userfile := "test.txt"
	fout,err := os.Create(userfile)
	defer fout.Close()
	
	if err != nil {
		fmt.Println(userfile,err)
	}

	files := make([]string, 0, 1000)
	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return err
	}

	//PthSep := string(os.PathSeparator)

	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { 
			continue
		}
		//if reg := regexp.MustCompile(`(?i:^hello)`){
		//	continue
		//}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			hashName,err := hashFile(fi.Name())
			if err != nil {
				return err
			}
			files = append(files,fmt.Sprintf("%s,%s,%d",fi.Name(),hashName,fi.Size()))
		}
		fout.WriteString("files\r\n")
		}
	return nil
}
