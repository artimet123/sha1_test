ckage main

import (
"fmt"
"io/ioutil"
"os"
"path/filepath"
"strings"
)

func ListDir(dirPth string, suffix string) (err error) {
	userfile := "test.txt"
	fout,err := os.Create(userfile)
	defer fout.close()
	
	if err != nil {
		fmt.Println(userfile,err)
	}

	files = make([]string, 0, 1000)
	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return err
	}

	PthSep := string(os.PathSeparator)

	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { 
			continue
		}
		if reg := regexp.MustCompile(`(`(?i:^seq)`)){
			continue
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
hashName,err := hashFile(fi)
			files = append(files, fi.Name()+','+hashName+','+fi.Size())
		}
		fout.WriteString("files\r\n")

	}
	return nil
}
