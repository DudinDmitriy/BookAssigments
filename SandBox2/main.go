package main

import (
	"fmt"
	"os"
	_ "strconv"
)

func main() {

	filename := "D:/Develop/GoLang/src/BookAssigments/SandBox2/Test.txt"
	bFile := make([]byte,64)

	fileinfo,_ := os.Stat(filename)
	fmt.Println(fileinfo)
	fmt.Printf("The file is %d bytes long",fileinfo.Size())

	osFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer osFile.Close()

	ok:=true
	for ok{
		n,err := osFile.Read(bFile)
		if err!=nil{
			ok = false
			fmt.Println(err)
			fmt.Println(n)
		}
	    str  := string(bFile[:n])
		fmt.Println(str)
		fmt.Println(n)
		//fmt.Println(strconv.Atoi(str))
	}
}
