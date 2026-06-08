package main

import (
	"fmt"
	"sync"
)
type Data struct {
	Val int
}

func main () {
ch := make(chan int, 5)
var wg sync.WaitGroup
var mu sync.Mutex
slice := make([]int,0,5)
	wg.Add(1)
	go Worker(1, ch, &wg, &slice, &mu)

dataTest := [5]int{3, 6, 19, 30 ,2 }
for _,v := range dataTest {
	ch <- v
}
close(ch)
wg.Wait()
fmt.Println(slice)
}

func Worker(id int , ch <- chan int, wg *sync.WaitGroup, slice *[]int, mu *sync.Mutex) {
	defer wg.Done()
	data := Data{}
	data.filterData( ch, slice, mu)
	mu.Lock()
	data.ProccessData(slice)
	mu.Unlock()
	} 

func (D Data)filterData(results <- chan int, slice *[]int, mu *sync.Mutex) {
	for v := range results {
		if (v % 2 == 0) {
			mu.Lock()

			*slice = append(*slice, v)
			fmt.Printf("Un worker agrego el valor %d \n", v)
			mu.Unlock()
		}
	}
}
func (D Data)ProccessData(slice *[]int) {
snapshot := append([]int{}, *slice...)
	
	for _, v := range snapshot {
		v = v * 10
		*slice = append(*slice, v)
	}
}