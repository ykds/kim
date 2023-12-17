package mq

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Host     string `json:"host" yaml:"host"`
	Post     string `json:"post" yaml:"post"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Vhost    string `json:"vhost" yaml:"vhost"`
}

type RabbitMQClient struct {
	*amqp091.Connection
}

func NewRabbitMQClient(cfg Config) *RabbitMQClient {
	if cfg.Vhost == "" {
		cfg.Vhost = "/"
	}
	dns := fmt.Sprintf("amqp://%s:%s@%s:%s%s", cfg.User, cfg.Password, cfg.Host, cfg.Post, cfg.Vhost)
	connection, err := amqp091.Dial(dns)
	if err != nil {
		panic(err)
	}
	return &RabbitMQClient{
		connection,
	}
}

func (r *RabbitMQClient) GetChannel() (*RabbitMQChannel, error) {
	channel, err := r.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQChannel{channel}, nil
}

func (r *RabbitMQClient) Close() error {
	return r.Close()
}

type option struct {
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	exclusive  bool
	args       map[string]interface{}
}

type Option func(option *option)

func WithDurable(durable bool) Option {
	return func(option *option) {
		option.durable = durable
	}
}

func WithAutoDelete(autoDelete bool) Option {
	return func(option *option) {
		option.autoDelete = autoDelete
	}
}

func WithInternal(internal bool) Option {
	return func(option *option) {
		option.internal = internal
	}
}

func WithNoWait(noWait bool) Option {
	return func(option *option) {
		option.noWait = noWait
	}
}

func WithExclusive(exclusive bool) Option {
	return func(option *option) {
		option.exclusive = exclusive
	}
}

func WithArgs(args map[string]interface{}) Option {
	return func(option *option) {
		option.args = args
	}
}

type RabbitMQChannel struct {
	*amqp091.Channel
}

func (rc *RabbitMQChannel) DeclareExchange(name string, kind string, opts ...Option) {
	opt := new(option)
	for _, o := range opts {
		o(opt)
	}
	err := rc.ExchangeDeclare(name, kind, opt.durable, opt.autoDelete, opt.autoDelete, opt.noWait, opt.args)
	if err != nil {
		panic(err)
	}
}

func (rc *RabbitMQChannel) DeclareQueue(name string, opts ...Option) {
	opt := new(option)
	for _, o := range opts {
		o(opt)
	}
	_, err := rc.QueueDeclare(name, opt.durable, opt.autoDelete, opt.exclusive, opt.noWait, opt.args)
	if err != nil {
		panic(err)
	}
}

func (rc *RabbitMQChannel) BindingQueue(exchangeName string, queueName string, routingKey string, opts ...Option) {
	opt := new(option)
	for _, o := range opts {
		o(opt)
	}
	err := rc.QueueBind(queueName, routingKey, exchangeName, opt.noWait, opt.args)
	if err != nil {
		panic(err)
	}
}

func (rc *RabbitMQChannel) PublishMsg(ctx context.Context, exchange, key string, msg []byte) error {
	return rc.PublishWithContext(ctx, exchange, key, false, false, amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 2,
		Body:         msg,
	})
}
