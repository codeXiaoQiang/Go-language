package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for{
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(10) +1)
	}

}
