package main

import (
	"fmt"
	"sync"
	"time"
)

func bell(b *sync.WaitGroup) {

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("in bell function")
	b.Done() //tells compiler thats done
}

func main() {
	var b sync.WaitGroup
	b.Add(1)
	fmt.Println("in main function")
	go bell(&b)
	b.Wait() //runs untill becomes 0
	fmt.Println("execution done")

}
