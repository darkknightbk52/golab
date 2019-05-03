package memory

import (
	"fmt"
	"github.com/darkknightbk52/golab/go-micro/broker"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"testing"
)

func Test(t *testing.T) {
	RegisterTestingT(t)

	const (
		topicHeader = "Topic"
		idHeader    = "ID"
	)

	type testCase struct {
		message *broker.Message
	}

	var testCases []testCase
	handledMessages := make(map[string][]*broker.Message)
	{
		msg := broker.NewMessage()
		msg.Header[topicHeader] = "Topic_1"
		msg.Header[idHeader] = uuid.New().String()
		msg.Body = []byte("LocTD")
		testCases = append(testCases, testCase{
			message: msg,
		})
	}

	{
		msg := broker.NewMessage()
		msg.Header[topicHeader] = "Topic_2"
		msg.Header[idHeader] = uuid.New().String()
		msg.Body = []byte("GiangPT")
		testCases = append(testCases, testCase{
			message: msg,
		})
	}

	b := NewMemoryBroker()
	b.Connect()
	defer b.Disconnect()

	handler := func(p broker.Publication) error {
		msg := p.Message()
		if msg == nil || msg.Header[topicHeader] == "" {
			err := errors.New("Invalid message")
			t.Error(err)
			return err

		}
		handledMessages[msg.Header[topicHeader]] = append(handledMessages[msg.Header[topicHeader]], p.Message())
		fmt.Println("Processed message:", p.Message())
		return nil
	}

	for _, c := range testCases {
		topic := c.message.Header[topicHeader]
		s, err := b.Subscribe(topic, handler)
		Expect(err).Should(Succeed())
		fmt.Println("Subscribed on topic:", topic)

		err = b.Publish(topic, *c.message)
		Expect(err).Should(Succeed())
		fmt.Println("Published message:", c.message)

		err = s.Unsubscribe()
		Expect(err).Should(Succeed())
	}

	for _, c := range testCases {
		msgs, ok := handledMessages[c.message.Header[topicHeader]]
		Expect(ok).Should(BeTrue())
		found := false
		for _, m := range msgs {
			if m.Header[idHeader] == c.message.Header[idHeader] {
				found = true
				Expect(m.Body).Should(Equal(c.message.Body))
			}
		}
		Expect(found).Should(BeTrue())
	}
}
