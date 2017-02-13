package main

import (
"fmt"
"io/ioutil"
"os"
"path/filepath"
"strings"
)

//��ȡָ��Ŀ¼�µ������ļ�����������һ��Ŀ¼����������ƥ���׺���ˡ�
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

	suffix = strings.ToUpper(suffix) //���Ժ�׺ƥ��Ĵ�Сд

	for _, fi := range dir {
		if fi.IsDir() { // ����Ŀ¼
		continue
		}
		//�жϸ�Ŀ¼�µ��ļ��ĺ�׺��suffix��β�Ƿ�ƥ��
		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //ƥ��
			continue
		}
	
		if reg := regexp.MustCompile(`(`(?i:^seq)`){
			continue
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //ƥ���ļ�
			hashName,err := hashFile(fi)
			files = append(files, fi.Name()+','+hashName+','+fi.Size())
		}
		fout.WriteString("files\r\n")

	}
	return nil
}