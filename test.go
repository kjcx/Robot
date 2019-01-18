package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var i int
	for i=0;i<20;i++ {
		go randnum()
	}
	time.Sleep(1000*time.Second)
}

func randnum(){
	for true  {
		n := rand.Intn(6)
		fmt.Println(n)
		time.Sleep(2*time.Second)
	}
}