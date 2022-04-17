package main

import (
	"bytes"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	"thoth/avro"
	"time"
)

var pcClient, err = NewPostgresClient(GetDatabaseTestConfig())

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	r := bytes.NewReader(msg.Payload())
	data, err := avro.DeserializeTelemetryDataBatch(r)
	fmt.Printf("Received message: \n%s\n error: %v\n", data.Size, err)
	err = pcClient.BatchInsertTelemetryData(&data)
	if err != nil {
		panic(err)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "127.0.0.1"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.CleanSession = false
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	time.Sleep(6000 * time.Second)

	client.Disconnect(250)
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
