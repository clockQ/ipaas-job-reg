package PhMessage

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"log"
)

type PhKafkaHandler struct {
	schemaRepositoryUrl string
	bkc                 *bmkafka.Config
}

func (handler PhKafkaHandler) New(srUrl string) *PhKafkaHandler {
	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err)
	}

	return &PhKafkaHandler{
		schemaRepositoryUrl: srUrl,
		bkc:                 bkc,
	}
}

func (handler PhKafkaHandler) Send(topic string, model PhModel.PhAvroModel) (err error) {
	record, err := model.GenSchema(model).GenRecord(model)
	if err != nil {
		return
	}
	log.Printf("Kafka 发送消息 %s 到 %s \n", record.String(), topic)

	encoder := kafkaAvro.NewKafkaAvroEncoder(handler.schemaRepositoryUrl)
	recordByteArr, err := encoder.Encode(record)
	if err != nil {
		return
	}

	handler.bkc.Produce(&topic, recordByteArr)
	return
}

func (handler PhKafkaHandler) Linster(topics []string, msgModel interface{}, subscribeFunc func(receive interface{})) {
	handler.bkc.SubscribeTopics(topics, func(receive interface{}) {
		decoder := kafkaAvro.NewKafkaAvroDecoder(handler.schemaRepositoryUrl)
		record, err := decoder.Decode(receive.([]byte))
		if err != nil {
			errMsg := fmt.Sprintf("Kafka 接受的 %s 信息解析出错: %s", topics, err)
			log.Println(errMsg)
			bmlog.StandardLogger().Error(errMsg)
			return
		}

		err = json.Unmarshal([]byte(record.(*avro.GenericRecord).String()), msgModel)
		if err != nil {
			errMsg := fmt.Sprintf("Kafka 接受的 %s 信息解析出错: %s", topics, err)
			log.Println(errMsg)
			bmlog.StandardLogger().Error(errMsg)
			return
		}

		log.Printf("Kafka 接受从 %s 来的消息 %#v \n", topics, msgModel)
		subscribeFunc(msgModel)
	})
}