package amqpextra_test

import (
	"log"

	"github.com/makasim/amqpextra"
	"github.com/makasim/amqpextra/publisher"
	"github.com/streadway/amqp"
)

func ExampleDialer_Publisher() {
	// open connection
	dialer, err := amqpextra.NewDialer(amqpextra.WithURL("amqp://guest:guest@localhost:5672/%2f"))
	if err != nil {
		log.Fatal(err)
	}

	// create publisher
	p, err := dialer.Publisher()
	if err != nil {
		log.Fatal(err)
	}

	// publish a message
	go p.Publish(publisher.Message{
		Key: "test_queue",
		Publishing: amqp.Publishing{
			Body: []byte(`{"foo": "fooVal"}`),
		},
	})

	// close publisher
	p.Close()
	<-p.NotifyClosed()

	// close connection
	dialer.Close()

	// Output:
}

func ExampleNewPublisher() {
	// you can get readyCh from dialer.ConnectionCh() method
	var connCh chan *amqpextra.Connection

	// create publisher
	p, err := amqpextra.NewPublisher(connCh)
	if err != nil {
		log.Fatal(err)
	}

	// publish a message
	go p.Publish(publisher.Message{
		Key: "test_queue",
		Publishing: amqp.Publishing{
			Body: []byte(`{"foo": "fooVal"}`),
		},
	})

	// close publisher
	p.Close()
	<-p.NotifyClosed()

	// Output:
}
