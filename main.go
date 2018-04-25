package main

import (
	"github.com/irvana/kstream/producer"
	"github.com/irvana/kstream/consumer"
)

func main() {
	producer.RunProducer()
	consumer.RunConsumer()
}