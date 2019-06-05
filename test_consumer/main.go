package main

import (
	"fmt"
	"myself/consumer"
	"time"

	"github.com/nsqio/go-nsq"
)

func handle(msg *nsq.Message) bool {
	fmt.Printf("msg.msg=%v,  msg=%s \n", time.Unix(0, msg.Timestamp).Format("2006-01-02 03:04:05"), string(msg.Body))
	return true
}

func main() {
	handle222 := consumer.ConsumerHandle{
		Cb: handle,
	}

	consumer := consumer.Consumer{
		TopicName:      "test",
		Channel:        "channel",
		NsqLookupdAddr: "127.0.0.1:4161", //方式1：连接nsqlookup
		NsqAddr:        "127.0.0.1:4151", //方式2：直连nsq
		Handle:         handle222,
	}

	consumer.Init()
	consumer.Run()

	select {}
}
