package consumer

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type CallBackFunc func(msg *nsq.Message) bool

type ConsumerHandle struct {
	Cb CallBackFunc
}

func (h *ConsumerHandle) HandleMessage(msg *nsq.Message) error {
	h.Cb(msg)
	return nil
}

type Consumer struct {
	Handle         ConsumerHandle
	TopicName      string
	Channel        string
	NsqLookupdAddr string
	NsqAddr        string
}

func (c *Consumer) Init() error {
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second
	config.MaxInFlight = 2//大于等于实际的nsqd数量
	consumer, err := nsq.NewConsumer(c.TopicName, c.Channel, config)
	if err != nil {
		panic(err)
	}

	consumer.SetLogger(nil, 0)

	handle := ConsumerHandle{
		Cb: c.Handle.Cb,
	}
	consumer.AddHandler(&handle)

	//连接nsqlookup，自动发现nsq。
	err = consumer.ConnectToNSQLookupd(c.NsqLookupdAddr)
	if err != nil {
		fmt.Println(err)
	}

	//直连nsq
	// err = consumer.ConnectToNSQD(c.NsqAddr)
	// if err != nil {
	// 	fmt.Println("consumer.ConnectToNSQD error...")
	// }

	fmt.Println("consumer.ConnectToNSQD ok ..")
	return nil
}

func (c *Consumer) Run() {
	for range time.NewTicker(time.Second * 10).C {
		fmt.Printf("Consumer is run   %v...\n", time.Now().Format("2006-01-02 03:04:05"))
	}
}
