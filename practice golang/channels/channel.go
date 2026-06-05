package main

import (
	"fmt"
	"time"
)

func BaristAtent(coffesReady chan string) {
i := 0

for  pedido := range coffesReady {
	fmt.Printf("En preparacion cafe %v, %v \n", i , pedido)
	time.Sleep(1 * time.Second)
	i++
}
}


func main() {
	cafeReadys := make(chan string)
	go BaristAtent(cafeReadys)

	for i:= 0; i < 10; i++ {
		pedido := fmt.Sprintf("Cafe #%d \n", i)
		fmt.Printf("[Cliente %d]: Intentando dejar mi pedido en la barra\n", i)

		cafeReadys <- pedido
	}
}

