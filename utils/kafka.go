package utils

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

//Brokers:  []string{brokerAddr},
//GroupID:  groupId, // 指定消费者组id
//Topic:    topic,
//MaxBytes: 10e6, // 10MB

//Addr:     kafka.TCP(brokerAddr),
//Topic:    topic,
//Balancer: &kafka.Hash{},

const (
	Topic      = "message"
	MaxBytes   = 10e6
	brokerAddr = "localhost:9092"
)

type Kafka struct {
	r *kafka.Reader
	w *kafka.Writer
}

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

func KafkaInit() *Kafka {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddr},
		GroupID:  Topic, // 指定消费者组id
		Topic:    Topic,
		MaxBytes: MaxBytes,
	})
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddr),
		Topic:    Topic,
		Balancer: &kafka.Hash{},
	}
	kaf := &Kafka{r: reader, w: writer}
	return kaf
}

func (k *Kafka) Writer(sender, input, id string) error {
	msg := Message{
		Receiver: id,
		Sender:   sender,
		Content:  input,
	}
	jsonData, _ := json.Marshal(msg)
	err := k.w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(sender),
		Value: jsonData,
	})
	if err != nil {
		k.w.Close()
		return err
	}
	return nil
}

func (k *Kafka) Close() {
	if k.r != nil {
		k.r.Close()
	}
	if k.w != nil {
		k.w.Close()
	}
}

func (k *Kafka) Reader(ctx context.Context, ch chan Message) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka Reader 退出")
			return nil
		default:
			m, err := k.r.ReadMessage(ctx)
			if err != nil {
				log.Println("读取 Kafka 消息失败：", err)
				return err
			}

			var msg Message
			err = json.Unmarshal(m.Value, &msg)
			if err != nil {
				log.Println("消息反序列化失败：", err)
				return err
			}
			ch <- msg
		}
	}
}
