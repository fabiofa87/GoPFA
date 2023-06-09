package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fabiofa8/pfa-go/internal/order/infra/database"
	"github.com/fabiofa8/pfa-go/internal/order/usecase"
	"github.com/fabiofa8/pfa-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)


func main() {
	maxWorkers := 4
	wg := sync.WaitGroup{}
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	deliveryMessage := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, deliveryMessage)
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go worker(deliveryMessage, uc, i)
		defer wg.Done()
	}
	wg.Wait()
}

func worker(deliveryMessage <- chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println(err)
		}
		input.Tax = 10.2

		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println(err)
		}

		msg.Ack(false)
		
	}
}