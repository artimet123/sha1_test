package main

import (
"fmt"
"io/ioutil"
"os"
"strings"
"runtime"
"path"
)

func ListDir(dirPth string, suffix string)(err error) {
	userfile := "test.txt"
	fout,err := os.Create(userfile)
	defer fout.Close()
	
	if err != nil {
		fmt.Println(userfile,err)
	}

	//files := make([]string, 0, 1000)
	var shaString string = ""
	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return err
	}

	//PthSep := string(os.PathSeparator)

	//suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		fmt.Printf("mao1111-%s",fi.Name())
	_, fulleFilename, line, _ := runtime.Caller(0)
	fmt.Println(fulleFilename)
	fmt.Println(line)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)
    	fmt.Println("filenameWithSuffix=", filenameWithSuffix)
    	var fileSuffix string
    	fileSuffix = path.Ext(filenameWithSuffix)
    	fmt.Println("fileSuffix=", fileSuffix)
     
    var filenameOnly string
    filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
    fmt.Println("filenameOnly=", filenameOnly)
		if strings.EqualFold(fileSuffix, suffix){
			continue
		}

		//if reg := regexp.MustCompile(`(?i:^hello)`){
		//	continue
		//}
		//if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			hashName,err := hashFile(fi.Name())
			if err != nil {
				return err
			}
			shaString = fmt.Sprintf("%s,%s,%d\n",fi.Name(),hashName,fi.Size())
			//files = append(files,fmt.Sprintf("%s,%s,%d\n",fi.Name(),hashName,fi.Size()))
		//}
		fout.WriteString(shaString)
		}
	return nil
}
