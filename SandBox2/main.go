package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// func main() {

// 	filename := "D:/Develop/GoLang/src/BookAssigments/SandBox2/Test.txt"
// 	bFile := make([]byte,64)

// 	fileinfo,_ := os.Stat(filename)
// 	fmt.Println(fileinfo)
// 	fmt.Printf("The file is %d bytes long",fileinfo.Size())

// 	osFile, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer osFile.Close()

// 	ok:=true
// 	for ok{
// 		n,err := osFile.Read(bFile)
// 		if err!=nil{
// 			ok = false
// 			fmt.Println(err)
// 			fmt.Println(n)
// 		}
// 	    str  := string(bFile[:n])
// 		fmt.Println(str)
// 		fmt.Println(n)
// 		//fmt.Println(strconv.Atoi(str))
// 	}
// }

func main() {

	filepath := "D:/Develop/GoLang/src/GoLangTest/dictionary/Data/prod.csv"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readcsv := csv.NewReader(file)
	ok := true
	for ok {
		rec, err := readcsv.Read()
		if err != nil {
			break
		}
		for _; i:=range rec;{
			fmt.Printf("Значение поля: %s", i)
		}
	}

	// var datafile string
	// var bytefile []byte

	// var strBuilder strings.Builder
	// bytefile = make([]byte, 64)

	// ok := true
	// for ok {
	// 	n, err := file.Read(bytefile)
	// 	if err != nil {
	// 		ok = false
	// 	}
	// 	file.Read(bytefile)
	// 	strBuilder.Write(bytefile[:n])

	// }
	// strres := strBuilder.String()
	// readcsv := csv.NewReader(file)
	// readcsv.Comma =";"
}
