package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	<-time.After(15 * time.Second)
	fmt.Println(time.Now().Sub(t).String())
	fmt.Println(t.Sub(time.Now()).String())
}
