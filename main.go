package main

import (
	"fmt"
	"os"
	"time"

	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var sendIt = false

var onMessageReceived mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Distance: %s\n", msg.Payload())
	dist, err := strconv.ParseInt(string(msg.Payload()), 10, 32)
	if err != nil {
		fmt.Println("Error converting")
	}

	if dist < 80 && !sendIt {
		sendEmail()
		sendIt = true
	}
}

func main() {

	const TOPIC = "/home/cave/sensors/distance"

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.2.99:1883")
	opts.SetDefaultPublishHandler(onMessageReceived)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Println("Connected to server...")
	}

	if token := client.Subscribe(TOPIC, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		time.Sleep(1 * time.Second)
		if sendIt {
			timer := time.NewTimer(time.Minute * 2)
			go func() {
				<-timer.C
				sendIt = false
			}()
		}
	}
}

func sendEmail() {
	sender := NewSender("homeappmqtt@gmail.com", os.Getenv("HOMEAPPPASSW"))
	Receiver := []string{"agarcia@cittec.es"}

	Subject := "Distance Sensor Activated!"
	bodyMessage := "It seems someone is in..."

	sender.SendMail(Receiver, Subject, bodyMessage)
}
