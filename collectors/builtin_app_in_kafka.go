package collectors

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/tidwall/gjson"
)

type AppConnection struct {
	kafkaconn         []string
	OutChan           *chan *CollectData
	Exit              chan int
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

func GetAppConn(address []string, Out *chan *CollectData) (conn *AppConnection, err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		return conn, err
	}
	partitionConsumer, err := consumer.ConsumePartition("dw30_metrics", 0, sarama.OffsetNewest)
	if err != nil {
		return conn, err
	}
	return &AppConnection{
		kafkaconn:         address,
		Exit:              make(chan int),
		OutChan:           Out,
		consumer:          consumer,
		partitionConsumer: partitionConsumer,
	}, err
}

func (app *AppConnection) loadStatus() {

}

func (app *AppConnection) Start() {
	consumed := 0
ConsumerLoop:
	for {
		select {
		case <-app.Exit:
			break ConsumerLoop
		case msg := <-app.partitionConsumer.Messages():
			//log.Printf("Consumed message offset %d\n", msg.Offset)
			//log.Printf("Consumed message offset %s\n", msg.Value)
			result := gjson.Get(string(msg.Value), "instance.app")
			//log.Printf("app: %s\n", result.String())
			switch result.String() {
			case "munich":
				var munichapp Munich
				json.Unmarshal(msg.Value, &munichapp)
				munichapp.toCollectData(app.OutChan)
			case "eau-rouge":
				var eaurougeapp EauRouge
				json.Unmarshal(msg.Value, &eaurougeapp)
				eaurougeapp.toCollectData(app.OutChan)
			case "buy-admin":
				var buyadminapp BuyAdmin
				json.Unmarshal(msg.Value, &buyadminapp)
				buyadminapp.toCollectData(app.OutChan)
			case "buy":
				var buyapp Buy
				json.Unmarshal(msg.Value, &buyapp)
				buyapp.toCollectData(app.OutChan)
			case "saturn":
				var saturnapp Saturn
				json.Unmarshal(msg.Value, &saturnapp)
				saturnapp.toCollectData(app.OutChan)
			case "vienna":
				var viennaapp Vienna
				json.Unmarshal(msg.Value, &viennaapp)
				viennaapp.toCollectData(app.OutChan)
			case "mandala":
				var mandalaapp Mandala
				json.Unmarshal(msg.Value, &mandalaapp)
				mandalaapp.toCollectData(app.OutChan)
			case "jerusalem":
				var jerusalemapp Jerusalem
				json.Unmarshal(msg.Value, &jerusalemapp)
				jerusalemapp.toCollectData(app.OutChan)
			default:
				log.Printf("app: %s not support \n", result.String())
			}
			consumed++
		}
	}
}

func (app *AppConnection) Close() {
	close(app.Exit)
}
