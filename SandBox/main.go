package main

import "math/rand"

//import "time"
import "fmt"

//import "sync"

//var wg sync.WaitGroup

// func generateint(ch chan int, r int) {
// 	var i int
// 	for i = 1; i < r; i++ {
// 		newr := rand.Intn(i)
// 		fmt.Println("Генерация значения ", newr)
// 		ch <- newr
// 		//time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }

func generateint(ch chan<- int, r int) {
	//defer wg.Done()
	newr := rand.Intn(r)
	fmt.Println("Генерация значения ", newr)
	ch <- newr
	//time.Sleep(100 * time.Second)
	if newr > 1 {
		//wg.Add(1)
		go generateint(ch, newr)
	} else {
		close(ch)
	}
}

func printint(ch <-chan int, chEnd chan<- int, r int) {

	var ii int
	ok := true
	//	for i := 1; i < (r); i++ {
	for ok {
		ii, ok = <-ch
		fmt.Println("Вывод значения ", ii, ok)
		//time.Sleep(1 * time.Second)
	}
	//	wg.Done()
	close(chEnd)
}

func main() {

	const n = 100
	ch := make(chan int, 1000000)
	chEnd := make(chan int, 0)
	//wg.Add(1)
	go generateint(ch, n)
	//fmt.Println("Вывод значения(главный) ", <-ch)
	//wg.Add(1)
	go printint(ch, chEnd, n)
	//wg.Wait()
	<-chEnd
}
