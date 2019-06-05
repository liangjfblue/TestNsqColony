package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

var (
	tcpNsqdAddrr = "127.0.0.1:4151"
)

func main() {
	config := nsq.NewConfig()

	tPro, err := nsq.NewProducer(tcpNsqdAddrr, config)
	if err != nil {
		fmt.Println(err)
	}
	topic := "test"
	tCommand := "test 4151"

	err = tPro.Publish(topic, []byte(tCommand))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("publish msg ok")
}
