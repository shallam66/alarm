package kafka

import (
	"fmt"
	sarama "github.com/Shopify/sarama"
)

//type KafkaClient interface {

//func CreateClient()

func Consume() {
	config := sarama.NewConfig()
	//config.ChannelBufferSize =
	master, _ := sarama.NewClient([]string{"localhost:9095"}, config)
	cmas, _ := sarama.NewConsumerFromClient(master)
	offsetmanager, _ := sarama.NewOffsetManagerFromClient("test2", master)
	part_man, _ := offsetmanager.ManagePartition("tt.test", 1)
	//part_man.MarkOffset(0, "tt.test")
	off1, mata1 := part_man.NextOffset()
	fmt.Println(off1)
	fmt.Println(mata1)
	csm, _ := cmas.ConsumePartition("tt.test", 1, off1)
	//csm.Messages
	//Loop:
	for i := 0; i < 200; i++ {
		select {
		case message := <-csm.Messages():
			off1++
			part_man.MarkOffset(off1, "tt.test")
			fmt.Println(message)
		case err := <-csm.Errors():
			fmt.Println(err)
			//csm.Close()
			//			break Loop
		}
	}
	part_man.MarkOffset(off1, "tt.test")
	csm.Close()
	fmt.Println(off1)
	part_man.Close()
	//part_man2, _ := offsetmanager.ManagePartition("syslog-ng", 0)
	//off2, mata2 := part_man2.NextOffset()
	//part_man.MarkOffset(off2, "syslog-ng")
	//fmt.Println(off2)
	//fmt.Println(mata2)
	cmas.Close()
	master.Close()
}
