package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

// Sarama configuration options
var (
	brokers  = "localhost:9092"
	version  = "2.8.1"
	group    = "test_group"
	topics   = "test"
	assignor = "sticky"
	oldest   = true
	verbose  = false
)

//
//func init() {
//	flag.StringVar(&brokers, "brokers", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
//	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
//	flag.StringVar(&version, "version", "2.1.1", "Kafka cluster version")
//	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma separated list")
//	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
//	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
//	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
//	flag.Parse()
//
//	if len(brokers) == 0 {
//		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
//	}
//
//	if len(topics) == 0 {
//		panic("no topics given to be consumed, please set the -topics flag")
//	}
//
//	if len(group) == 0 {
//		panic("no Kafka consumer group defined, please set the -group flag")
//	}
//}

func start() {
	log.Println("Starting a new Sarama consumer")
	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ready: make(chan struct{}),
	}

	ctx, cancel := context.WithCancel(context.Background())

	// 创建client
	client, err := newClient()
	if err != nil {
		panic(err)
	}

	// 根据client创建consumerGroup
	consumerGroup, err := sarama.NewConsumerGroupFromClient(group, client)
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumerGroup.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				select {
				case <-consumer.ready: //有成功消费调用过Setup，通道已被关闭，上下文取消进行退出
					return
				default: //进行消费失败后未执行Setup，上下文取消进行退出 需关闭通道
					close(consumer.ready)
					return
				}
			}
			select {
			case <-consumer.ready: //Setup有成功启动过，通道已被关闭
				consumer.ready = make(chan struct{})
			default:

			}
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func newClient() (sarama.Client, error) {
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	// 创建client
	client, err := sarama.NewClient(strings.Split(brokers, ","), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Partitions(client sarama.Client) ([]int32, error) {
	if client == nil {
		// 创建client
		var err error
		if client, err = newClient(); err != nil {
			panic(err)
		}
	}
	return client.Partitions(topics)
}

func Offsets(client sarama.Client, time int64) {
	if client == nil {
		// 创建client
		var err error
		if client, err = newClient(); err != nil {
			panic(err)
		}
	}
	partitions, _ := Partitions(client)
	for _, p := range partitions {
		offset, _ := client.GetOffset(topics, p, time)
		fmt.Println(fmt.Sprintf("%v %v %v", topics, p, offset))
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan struct{}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
