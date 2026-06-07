package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(idWorker int, colaImagenes <- chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for imagen := range colaImagenes {
		fmt.Printf("Worker #%d esta procesando imagen: %s  \n", idWorker, imagen)
		time.Sleep(2* time.Second)
		fmt.Printf("Worker #%d ha finalizado procesamiento de imagen: %s \n", idWorker, imagen)
	}
}

func main() {
	colaImagenes := make(chan string, 5) // tama;o del buffer de 5
	var wg sync.WaitGroup


	//creamos los workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(i, colaImagenes, &wg)
	}

	// enviar imagenes a la cola del buffer

	for i := 0; i < 10; i++ {
		imagen := fmt.Sprintf("Imagen #%d", i)
		fmt.Printf("Enviando la imagen #%d \n", i)
		colaImagenes <- imagen
	}
	close(colaImagenes)
	wg.Wait()
}
