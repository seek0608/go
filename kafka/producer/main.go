package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//消息写入kafka
func main() {
	//初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	client, err := sarama.NewSyncProducer([]string{"118.126.89.12:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}
	defer client.Close()

	for i := 0; i < 5; i++ {
		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "cctv1"
		msg.Value = sarama.StringEncoder("this is a good test,hello kai")
		//发送消息
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}
}
