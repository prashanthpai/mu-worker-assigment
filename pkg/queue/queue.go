package queue

import (
	"context"
	"errors"

	"github.com/streadway/amqp"
)

var ErrQueueClosed = errors.New("queue closed")

type Config struct {
	URI          string `toml:"uri" validate:"required" env:"QUEUE_URI,overwrite"`
	QueueName    string `toml:"queue_name" validate:"required" env:"QUEUE_NAME,overwrite"`
	ConsumerName string `toml:"consumer_name" validate:"required" env:"QUEUE_CONSUMER_NAME,overwrite"`
}

type Client struct {
	q            <-chan amqp.Delivery
	conn         *amqp.Connection
	ch           *amqp.Channel
	consumerName string
}

func New(config Config) (*Client, error) {
	// dial and create a connection

	// create a channel

	// decrare a qeueue

	// create a consumer (go channel) and store it in struct

	return nil, nil
}

func (c *Client) Close() {
	// cancel channel
	// close channel
	// close connection
}

func (c *Client) ReceiveMessage(ctx context.Context) ([]byte, error) {
	// receive message from go channel and return message content
	// return error if queue is closed
	return nil, nil
}
