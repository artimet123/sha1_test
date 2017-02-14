package main
import (
"os"
"fmt"
)

func main() {
	 HOME:= os.Getenv("HOME")
	fmt.Println(HOME)
	 err := ListDir("/home/cos/",".txt")
	if err != nil {
		fmt.Printf("error:%v",err)
		return
	}
}
