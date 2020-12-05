package main

import (
	"log"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second)

	<-t1.C
	log.Print("Timer 1 fired")

	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		log.Print("Timer 2 fired")
	}()
	stop2 := t2.Stop()
	if stop2 {
		log.Print("Timer 2 stopped")
	}

	time.Sleep(time.Second)
}
