package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Сейчас")
	Sleep(2)
	fmt.Println("Через 2 секунды")
}
