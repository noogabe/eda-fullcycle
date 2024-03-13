package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/noogabe/eda-fullcycle/consumer/internal/database"
	get_balance "github.com.br/noogabe/eda-fullcycle/consumer/internal/usecase/get"
	save_balance "github.com.br/noogabe/eda-fullcycle/consumer/internal/usecase/save"
	"github.com.br/noogabe/eda-fullcycle/consumer/internal/web"
	"github.com.br/noogabe/eda-fullcycle/consumer/internal/web/webserver"
	"github.com.br/noogabe/eda-fullcycle/consumer/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	getBalanceUsecase := get_balance.NewGetBalanceUsecase(database.NewBalanceDb(db))
	saveBalanceUsecase := save_balance.NewSaveBalanceUsecase(database.NewBalanceDb(db))

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	}

	topics := []string{"balances"}

	consumer := kafka.NewConsumer(&configMap, topics)

	consumerMsgChan := make(chan *ckafka.Message)

	go consumer.Consume(consumerMsgChan)

	consumerFunc := func(msgChan chan *ckafka.Message, saveBalanceUseCase *save_balance.SaveBalanceUsecase) {
		for {
			fmt.Println("Waiting for messages from kafka...")
			msg := <-msgChan
			kafkaMsg := KafkaMsgDto{}
			err := json.Unmarshal(msg.Value, &kafkaMsg)
			kafkaPayload := kafkaMsg.Payload
			if err != nil {
				fmt.Println(err.Error())
			}
			input := save_balance.SaveBalanceInputDto{
				AccountId: kafkaPayload.AccountIDFrom,
				Amount:    int(kafkaPayload.BalanceAccountIDFrom),
			}
			_, err = saveBalanceUseCase.Execute(input)
			if err != nil {
				fmt.Println(err.Error())
			}
			input = save_balance.SaveBalanceInputDto{
				AccountId: kafkaPayload.AccountIDTo,
				Amount:    int(kafkaPayload.BalanceAccountIDTo),
			}
			_, err = saveBalanceUseCase.Execute(input)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	go consumerFunc(consumerMsgChan, saveBalanceUsecase)

	getBalanceHandler := web.NewWebGetBalanceHandler(*getBalanceUsecase)

	webserver := webserver.NewWebServer(":3003")
	webserver.AddHandler("/balances/{id}", getBalanceHandler.GetBalance)
	fmt.Println("Server is running...")
	webserver.Start()

}

type PayloadKafkaDto struct {
	AccountIDFrom        string  `json:"account_id_from"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type KafkaMsgDto struct {
	Name    string
	Payload PayloadKafkaDto
}

func (kafkaPayload *PayloadKafkaDto) String() string {
	return fmt.Sprintf("AccountIDFrom: %s, BalanceAccountIDFrom: %f, AccountIDTo: %s, BalanceAccountIDTo: %f", kafkaPayload.AccountIDFrom, kafkaPayload.BalanceAccountIDFrom, kafkaPayload.AccountIDTo, kafkaPayload.BalanceAccountIDTo)
}
