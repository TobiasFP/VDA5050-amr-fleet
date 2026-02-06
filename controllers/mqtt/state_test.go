package mqttstate

import (
	"encoding/json"
	"testing"
	"time"

	"TobiasFP/BotNana/models"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type publishCall struct {
	topic   string
	payload string
}

type fakePublisher struct {
	calls []publishCall
	err   error
}

func (f *fakePublisher) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	var payloadStr string
	switch v := payload.(type) {
	case []byte:
		payloadStr = string(v)
	default:
		serialized, _ := json.Marshal(v)
		payloadStr = string(serialized)
	}
	f.calls = append(f.calls, publishCall{topic: topic, payload: payloadStr})
	return &fakeToken{err: f.err}
}

type fakeToken struct {
	err error
}

func (t *fakeToken) Wait() bool {
	return true
}

func (t *fakeToken) WaitTimeout(_ time.Duration) bool {
	return true
}

func (t *fakeToken) Done() <-chan struct{} {
	done := make(chan struct{})
	close(done)
	return done
}

func (t *fakeToken) Error() error {
	return t.err
}

func TestAssignOrderPublishesToOrderTopic(t *testing.T) {
	publisher := &fakePublisher{}
	order := models.Order{
		OrderID:      "order-123",
		Manufacturer: "acme",
		SerialNumber: "amr-1",
	}

	if err := AssignOrder(publisher, order); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(publisher.calls) != 1 {
		t.Fatalf("expected 1 publish call, got %d", len(publisher.calls))
	}
	call := publisher.calls[0]
	if call.topic != "vda5050/acme/amr-1/order" {
		t.Fatalf("expected topic vda5050/acme/amr-1/order, got %s", call.topic)
	}
	if call.payload == "" {
		t.Fatalf("expected payload to be set")
	}
}

func TestPublishInstantActionPublishes(t *testing.T) {
	publisher := &fakePublisher{}
	action := models.InstantAction{
		HeaderID:     1,
		Manufacturer: "acme",
		SerialNumber: "amr-1",
		Actions:      []models.Action{{ActionID: "a1", ActionType: "beep"}},
	}

	err := PublishInstantAction(publisher, action)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(publisher.calls) != 1 {
		t.Fatalf("expected 1 publish call, got %d", len(publisher.calls))
	}

	if publisher.calls[0].topic != "vda5050/acme/amr-1/instantActions" {
		t.Fatalf("expected topic vda5050/acme/amr-1/instantActions, got %s", publisher.calls[0].topic)
	}
}
