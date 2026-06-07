package main

import (
	"context"
	"fmt"
	"time"
)

func Worker( id int, messagesToProcess <- chan string, ctx context.Context) {
		for {

			select {
			case <- ctx.Done():
				fmt.Println("Message cancel for timeout")
				return
			case message, ok := <-messagesToProcess:
				if !ok{
					fmt.Printf("Worker %d, ha terminado, no hay mas mensajes \n", id)
					return
				}
				fmt.Printf("Worker %d, procesando %v \n",id , message)
				time.Sleep(2 * time.Second)
			}
			
			
		}
		}

func main() {
	ch := make(chan string, 10)
	ctx, cancel := context.WithTimeout(context.Background(), (5 * time.Second))
	defer cancel()
	go Worker(1, ch, ctx)
	go Worker(2, ch, ctx)
	go Worker(3, ch, ctx)

	go func () {for i:=0; i < 100; i++ {
		message := fmt.Sprintf("Mensaje #%d \n", i )
		ch <- message
		
	}
}()
time.Sleep(10 * time.Second)
}