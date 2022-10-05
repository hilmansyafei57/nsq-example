package nsq_fire_forget

import (
	"log"

	"github.com/nsqio/go-nsq"

	platform "github.com/firdasafridi/example-nsq"
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
	log.Println("Chanel NSQ Name", nff.env.Nsq)
	consum, err := nsq.NewConsumer(chanelName, "ch", config)
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
	log.Printf("Got a message: %v\n", message)
	log.Printf("The message %s\n", string(message.Body))
	return nil
}
