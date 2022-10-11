package nsq_producer

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nsqio/go-nsq"

	platform "github.com/hilmansyafei57/nsq-example"
)

type NsqProducer struct {
	producer *nsq.Producer
	env      *platform.Environment
}

func New() (np *NsqProducer, err error) {
	np = &NsqProducer{}

	np.env = platform.NewEnvironment()

	err = np.newProducer()
	if err != nil {
		log.Fatalln("Failed to create producer:", err)
		return nil, err
	}

	return np, nil
}

func (np *NsqProducer) Start() (err error) {
	log.Println("Start example nsq producer")
	np.CmdHandler()
	return
}

func (np *NsqProducer) Stop() (err error) {
	log.Println("Stop example nsq producer")
	np.producer.Stop()
	return
}

func (np *NsqProducer) newProducer() (err error) {
	config := nsq.NewConfig()
	np.producer, err = nsq.NewProducer(np.env.Nsq.Host, config)
	if err != nil {
		log.Fatalln("internal/pkg/nsq-producer/NsqProducer/newProducer", err)
	}

	log.Println("Success to create new producer")
	return nil
}

func (np *NsqProducer) Publish(b []byte) (err error) {
	err = np.producer.Publish(np.env.Nsq.ChanelName, b)
	if err != nil {
		return err
	}
	return nil
}

func (np *NsqProducer) CmdHandler() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 1000; i++ {
		np.Publish([]byte("test"))
	}

	for scanner.Scan() {
		command := scanner.Text()

		switch command {
		case "":
			continue
		case "exit":
			log.Println("Exit")
			os.Exit(1)
		default:
			err := np.Publish([]byte(command))
			if err != nil {
				log.Println("internal/pkg/nsq-producer/NsqProducer/CmdHandler:", err)
				continue
			}
			fmt.Println("Send:", command)
		}
	}
}
