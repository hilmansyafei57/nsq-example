package nsq_fire_forget

import (
	"log"

	"github.com/nsqio/go-nsq"

	platform "github.com/hilmansyafei57/nsq-example"
)

type NsqFireForget struct {
	consumer *nsq.Consumer
	env      *platform.Environment
}

func (nff *NsqFireForget) Start() (err error) {
	log.Println("Start example nsq fire n forget consumer")
	return nil
}

func (nff *NsqFireForget) Stop() {
	log.Println("Stop example nsq fire n forget")
}

func New() (nff *NsqFireForget, err error) {
	nff = &NsqFireForget{}

	nff.env = platform.NewEnvironment()

	err = nff.NsqHandler()
	if err != nil {
		log.Fatalln(err)
	}
	return nff, nil
}

func (nff *NsqFireForget) NsqHandler() (err error) {
	config := nsq.NewConfig()
	chanelName := nff.env.Nsq.ChanelName
	topicName := nff.env.Nsq.TopicName
	log.Println("Chanel NSQ Name", nff.env.Nsq)
	consum, err := nsq.NewConsumer(topicName, chanelName, config)
	if err != nil {
		return err
	}

	consum.AddHandler(nsq.HandlerFunc(message))
	err = consum.ConnectToNSQD(nff.env.Nsq.Host)
	if err != nil {
		return err
	}
	return nil
}

func message(message *nsq.Message) error {
	log.Printf("Receive message %s\n", string(message.Body))
	message.Finish()
	return nil
}
