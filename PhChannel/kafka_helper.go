package PhChannel

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
)

type PhKafkaHelper struct {
	schemaRepositoryUrl string
	bkc                 *bmkafka.Config
}

func (handler PhKafkaHelper) New(srUrl string) *PhKafkaHelper {
	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		panic(err)
	}

	return &PhKafkaHelper{
		schemaRepositoryUrl: srUrl,
		bkc:                 bkc,
	}
}

func (handler PhKafkaHelper) Send(topic string, model PhModel.PhAvroModel) (err error) {
	record, err := model.GenSchema(model).GenRecord(model)
	if err != nil {
		return
	}
	log.NewLogicLoggerBuilder().Build().Infof("Kafka 发送消息 %s 到 %s \n", record.String(), topic)

	encoder := kafkaAvro.NewKafkaAvroEncoder(handler.schemaRepositoryUrl)
	recordByteArr, err := encoder.Encode(record)
	if err != nil {
		return
	}

	handler.bkc.Produce(&topic, recordByteArr)
	return
}

func (handler PhKafkaHelper) Linster(topics []string, msgModel interface{}, subscribeFunc func(receive interface{})) {
	handler.bkc.SubscribeTopics(topics, func(receive interface{}) {
		decoder := kafkaAvro.NewKafkaAvroDecoder(handler.schemaRepositoryUrl)
		record, err := decoder.Decode(receive.([]byte))
		if err != nil {
			errMsg := fmt.Sprintf("Kafka 接受的 %s 信息解析出错: %s", topics, err)
			log.NewLogicLoggerBuilder().Build().Error(errMsg)
			return
		}

		err = json.Unmarshal([]byte(record.(*avro.GenericRecord).String()), msgModel)
		if err != nil {
			errMsg := fmt.Sprintf("Kafka 接受的 %s 信息解析出错: %s", topics, err)
			log.NewLogicLoggerBuilder().Build().Error(errMsg)
			return
		}

		log.NewLogicLoggerBuilder().Build().Infof("Kafka 接受从 %s 来的消息 %#v \n", topics, msgModel)
		subscribeFunc(msgModel)
	})
}
