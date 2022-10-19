package main

import (
	"log"

	"github.com/clferracin/minientregas-simulator/application/kafka"
	"github.com/clferracin/minientregas-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

}
