package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func worker(id int, numberToProccess <- chan int, result  chan int,ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Message cancel for timeout")
			return
		case message, ok := <-numberToProccess:
			if !ok {
				fmt.Printf("Worker %d, finished to process \n", id)
				return
			}
			fmt.Printf("Worker %d, proccessing %v \n", id, message)
	
			result <- message
			time.Sleep(10 * time.Millisecond)
		}
	}
}


func main() {
	var array []int
	var wg sync.WaitGroup
	ch := make(chan int, 20)
	result := make(chan int, 10000)
	ctx, cancel := context.WithTimeout(context.Background(), (30 * time.Second))
	defer cancel()
	for i:= 0; i < 1000; i++ {
		wg.Add(1)
		go worker(i, ch, result, ctx, &wg)
	}

	go func ()  {
		for i := 0; i < 100; i++ {
			number := i
			ch <- number
		}
		close(ch)
	}()

	go func ()  {
		wg.Wait()
		close(result)
	}()

	for int := range(result) {
		array = append(array, int)
	}


	fmt.Println(array)
	bubbleSort(array)
	fmt.Println(array)
}

func bubbleSort(numbers []int)[]int {
	for i := 0; i < len(numbers) - 1; i++ {
	for j := 0; j < len(numbers) - i - 1; j++ {
		if numbers[j] > numbers[j + 1] {
			numbers[j], numbers[j + 1] = numbers[j+1], numbers[j]
		}
	  }
	}
	
	return numbers
}