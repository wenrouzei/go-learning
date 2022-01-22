package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestConsumerGroup(t *testing.T) {
	_, err := Partitions(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(Partitions(nil))
	Offsets(nil, sarama.OffsetOldest)
	Offsets(nil, sarama.OffsetNewest)
	start()
}
