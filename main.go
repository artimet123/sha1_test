package main
import (
"os"
"fmt"
)

func main() {
	files, err := ListDir("D:\mygo\src\github.com\golang\snappy", "seq?"".txt") //���������Ŀ¼����ģ����˺�׺��txt�������ļ�
	fmt.Println(files, err)
}