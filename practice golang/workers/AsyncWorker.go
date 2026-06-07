package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, messagesToProcess<- chan int, ctx context.Context,  wg *sync.WaitGroup) {
defer wg.Done()
	sum := 0
	for {
			select {
			case <- ctx.Done():
				fmt.Println("Message cancel for timeout")
				return
			case message, ok := <-messagesToProcess:
				if !ok{
					fmt.Printf("Worker %d, ha terminado, no hay mas mensajes,  total = %v \n", id, sum)
					return
				}
				fmt.Printf("Worker %d, procesando %v \n",id , message)
				sum += message
				time.Sleep(1 * time.Second)
			}
			
			
		}

}
func main() {
var wg sync.WaitGroup
ch := make(chan int, 10)
ctx, cancel := context.WithTimeout(context.Background(), (5 * time.Second))
defer cancel()
	for i := 0; i < 3 ; i++ {
	wg.Add(1)
	go worker(i, ch, ctx, &wg)
	
	} 
	go func () {
		for i:=0; i < 100; i++ {
			message := i 
			ch <- message
			
		}
		}()
		wg.Wait()
}
