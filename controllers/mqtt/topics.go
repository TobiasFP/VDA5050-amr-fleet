package mqttstate

import (
	"errors"
	"fmt"
	"strings"

	"TobiasFP/BotNana/config"
)

const (
	defaultMqttTopicPrefix = "vda5050"
	topicState             = "state"
	topicConnection        = "connection"
	topicOrder             = "order"
	topicInstantActions    = "instantActions"
	topicFactsheet         = "factsheet"
	topicVisualization     = "visualization"
)

func topicPrefix() string {
	conf := config.GetConfig()
	if conf == nil {
		return defaultMqttTopicPrefix
	}
	prefix := conf.GetString("mqttTopicPrefix")
	if prefix == "" {
		prefix = defaultMqttTopicPrefix
	}
	return strings.Trim(prefix, "/")
}

func topicFor(manufacturer, serial, topic string) (string, error) {
	if manufacturer == "" || serial == "" {
		return "", errors.New("manufacturer and serialNumber are required to build a VDA5050 topic")
	}
	return fmt.Sprintf("%s/%s/%s/%s", topicPrefix(), manufacturer, serial, topic), nil
}

func wildcardTopic(topic string) string {
	return fmt.Sprintf("%s/+/+/%s", topicPrefix(), topic)
}

func parseTopic(topic string) (manufacturer, serial, leaf string, ok bool) {
	prefix := topicPrefix()
	prefixParts := strings.Split(strings.Trim(prefix, "/"), "/")
	parts := strings.Split(strings.Trim(topic, "/"), "/")
	if len(parts) != len(prefixParts)+3 {
		return "", "", "", false
	}
	if strings.Join(parts[:len(prefixParts)], "/") != prefix {
		return "", "", "", false
	}
	manufacturer = parts[len(prefixParts)]
	serial = parts[len(prefixParts)+1]
	leaf = parts[len(prefixParts)+2]
	return manufacturer, serial, leaf, true
}

func WildcardStateTopic() string {
	return wildcardTopic(topicState)
}

func WildcardConnectionTopic() string {
	return wildcardTopic(topicConnection)
}

func OrderTopic(manufacturer, serial string) (string, error) {
	return topicFor(manufacturer, serial, topicOrder)
}

func InstantActionsTopic(manufacturer, serial string) (string, error) {
	return topicFor(manufacturer, serial, topicInstantActions)
}
