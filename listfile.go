package main

import (
"fmt"
"io/ioutil"
"os"
"path/filepath"
"strings"
)

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (err error) {
	userfile := "test.txt"
	fout,err := os.Create(userfile)
	defer fout.close()
	
	if err != nil {
		fmt.Println(userfile,err)
	}

	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return err
	}

	PthSep := string(os.PathSeparator)

	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
		continue
		}
		//判断该目录下的文件的后缀和suffix结尾是否匹配
		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配
			continue
		}
	
		if reg := regexp.MustCompile(`(`(?i:^seq)`){
			continue
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			hashName,err := hashFile(fi)
			files = append(files, fi.Name()+','+hashName+','+fi.Size())
		}
		fout.WriteString("files\r\n")

	}
	return nil
}