package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Job started")
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Currently processing ", i, "th operation")
	}
	fmt.Println("Job ended")

}
