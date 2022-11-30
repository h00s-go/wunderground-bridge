package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/h00s-go/wunderground-bridge/config"
)

type MQTT struct {
	client mqtt.Client
}

func NewMQTT(c *config.MQTT) *MQTT {
	m := new(MQTT)

	opts := mqtt.NewClientOptions().AddBroker(c.Broker)
	opts.SetClientID(c.ClientID)
	opts.SetUsername(c.Username)
	opts.SetPassword(c.Password)
	opts.SetPingTimeout(60 * time.Second)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(30 * time.Second)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		log.Println("Connection lost to MQTT broker: ", err.Error())
	})
	opts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		log.Println("Reconnecting to MQTT broker...")
	})
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		log.Println("Connected to MQTT broker.")
	})

	m.client = mqtt.NewClient(opts)
	go func(m *MQTT) {
		for {
			if token := m.client.Connect(); token.Wait() && token.Error() != nil {
				log.Println(token.Error())
			} else {
				break
			}
			time.Sleep(30 * time.Second)
		}
	}(m)

	return m
}

func (m *MQTT) Publish(topic string, message string) {
	go func() {
		if m.client.IsConnected() {
			if token := m.client.Publish(topic, 0, false, message); token.Wait() && token.Error() != nil {
				log.Println(token.Error())
			}
		}
	}()
}

func (m *MQTT) Close() {
	m.client.Disconnect(1000)
}
