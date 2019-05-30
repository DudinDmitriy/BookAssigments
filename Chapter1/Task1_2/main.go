package main

import (
	"fmt"
"os"
"bufio")

func main(){
	
	counts := make(map[string]int)
	
	input:=bufio.NewScanner(os.Stdin)
		for input.Scan() {
       counts[input.Text]++
	}
	
	for line, n:=range counts{
		fmt.Print(n, line)
	}
}