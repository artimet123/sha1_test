package main

import (
"fmt"
"io/ioutil"
"os"
"regexp"
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
		if m, _ := regexp.MatchString(".*txt", fi.Name()); m {
			continue
		}
		//var reg = regexp.MustCompile(`.*txt`)
		//var fileSuffix string
		//fmt.Printf("%q\n", reg.FindAllString(fi.Name(), -1))
		//_,err := reg.FindAllString(fi.Name(), -1)
		//if strings.EqualFold(fileSuffix, suffix){
		//	continue
		//}

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
