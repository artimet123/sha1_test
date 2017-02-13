package main
import (
"os"
"fmt"
)

func main() {
	files, err := ListDir("D:\mygo\src\github.com\golang\snappy", "seq?"".txt") //这是找这个目录下面的，过滤后缀是txt，所有文件
	fmt.Println(files, err)
}