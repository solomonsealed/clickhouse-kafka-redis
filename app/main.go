package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

var (
	ctx = context.Background()
)

type RegionalActivityRecord struct {
	IdentityId  string `json:"identity_id"`
	Country     string `json:"country"`
	EventsCount uint64 `json:"events_count"`
}

func main() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	redisAddr := os.Getenv("REDIS_ADDR")

	err := kafkaToRedis(kafkaBroker, redisAddr)
	if err != nil {
		log.Fatalf("Failed to consume from Kafka and write to Redis: %v", err)
	}
}

func kafkaToRedis(kafkaBroker, redisAddr string) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "clickhouse_topic",
		GroupID: "clickhouse_group",
	})
	defer reader.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})
	defer rdb.Close()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		var regionalActivityRecord RegionalActivityRecord
		err = json.Unmarshal(msg.Value, &regionalActivityRecord)
		if err != nil {
			return err
		}

		fmt.Println(regionalActivityRecord.IdentityId, regionalActivityRecord.Country, regionalActivityRecord.EventsCount)
		err = rdb.HSet(ctx, regionalActivityRecord.IdentityId, regionalActivityRecord.Country, regionalActivityRecord.EventsCount).Err()
		if err != nil {
			return err
		}
	}
}
