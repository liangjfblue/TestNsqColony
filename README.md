1、安装库
go get github.com/nsqio/go-nsq

2、执行./start.sh
启动nsq服务
启动nsqd服务
启动nsqdadmin服务

3、进入test_consumer目录，执行 go run main.go
启动消费者

4、进入product目录，执行 go run product1.go / go run product2.go
