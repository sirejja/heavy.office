package consumer_group

import (
	"go.uber.org/zap"
	"route256/libs/logger"

	"github.com/Shopify/sarama"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready               chan bool
	topicHandlerMapping map[string]func(msg *sarama.ConsumerMessage) error
}

// NewConsumer - constructor
func New(topicHandlerMapping map[string]func(msg *sarama.ConsumerMessage) error) Consumer {
	return Consumer{
		ready:               make(chan bool),
		topicHandlerMapping: topicHandlerMapping,
	}
}

func (consumer *Consumer) Ready() <-chan bool {
	return consumer.ready
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
	for {
		select {
		case message := <-claim.Messages():
			topicHandler, ok := consumer.topicHandlerMapping[message.Topic]
			if ok {
				if err := topicHandler(message); err != nil {
					logger.Info("Message claimed",
						zap.ByteString("value", message.Value),
						zap.Time("timestamp", message.Timestamp),
						zap.String("topic", message.Topic),
					)
				}
				session.MarkMessage(message, "")
			} else {
				logger.Warn("Unbinded topic recieved")
				session.MarkMessage(message, "")
			}

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (consumer *Consumer) ToggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		logger.Warn("Resuming consumption")
	} else {
		client.PauseAll()
		logger.Warn("Pausing consumption")
	}

	*isPaused = !*isPaused
}
