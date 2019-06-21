package main
import "fmt"

type Product struct{
	idProduct int
	amount []int
}

func main() {
 p1 := Product{1,[]int{1,34,56,34,23}}
 p2 := Product{1,[]int{1,34,56,34,23}}
 p3 := Product{1,[]int{1,34,56,34,23}}
 p4 := Product{1,[]int{1,34,56,34,23}}
 p5 := Product{1,[]int{1,34,56,34,23}}
 p6 := Product{1,[]int{1,34,56,34,23}}
 go func(){
	 for i:=1;i<10000;i++{
	  p6.idProduct = p1.idProduct + p5.idProduct
	}
  }()

go func(){
	for i:=1;i<10000;i++{
		p2.idProduct = p3.idProduct + p4.idProduct
	}
}()
go func(){
	for i:=1;i<10000;i++{
		p3.idProduct = p4.idProduct + p5.idProduct
	}
}()

go func(){
	p4.idProduct = p5.idProduct + p5.idProduct
}()

go func(){
	p6.idProduct = p6.idProduct + 1
}()

fmt.Println(p1)
fmt.Println(p2)
fmt.Println(p3)
fmt.Println(p4)
fmt.Println(p5)
fmt.Println(p6)
}