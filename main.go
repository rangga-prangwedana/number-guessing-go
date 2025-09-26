package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello test")
	fmt.Println("The time now is", time.Now().UnixNano())
}
