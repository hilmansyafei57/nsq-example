package main

import (
	"log"
	"os"
	"os/signal"

	nsqproducer "github.com/hilmansyafei57/nsq-example/internal/pkg/nsq-producer"
)

func main() {
	np, err := nsqproducer.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		errServer := np.Start()
		if errServer != nil {
			log.Fatal(errServer)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	np.Stop()
}
