package reciewer

import (
	"github.com/bytedance/sonic"
	"github.com/osamikoyo/hrm-notify/internal/data/models"
	"github.com/osamikoyo/hrm-notify/internal/sender"
	"github.com/osamikoyo/hrm-notify/pkg/config"
	"github.com/osamikoyo/hrm-notify/pkg/loger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Reciewer struct{
	Sender *sender.Sender
	Logger loger.Logger
	AmqpQue amqp.Queue
	AmqpChannel *amqp.Channel
}

func Init(cfg *config.Config) (*Reciewer, error){
	sender := sender.New(cfg)


	conn, err := amqp.Dial(cfg.RabbitMQ)
	if err != nil{
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
  		"message", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)
	return &Reciewer{
		AmqpQue: q,
		AmqpChannel: ch,
		Logger: loger.New(),
		Sender: sender,
	}, err
}

func (r *Reciewer) WaitMessages() {
	msgs, err := r.AmqpChannel.Consume(
		r.AmqpQue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )
	  if err != nil{
		r.Logger.Error().Err(err)
	  }
	  
	  var message models.Msg

	  var forever chan struct{}
	  
	  go func() {
		for d := range msgs {
			if err = sonic.Unmarshal(d.Body, &message);err != nil{
				r.Logger.Error().Err(err)
			}

			msg := models.NewMessage(
				message.From,
				message.To,
				message.CC,
				message.Subject,
				message.Body,
			)

			err = r.Sender.Send(msg)
			if err != nil{
				r.Logger.Error().Err(err)
			}
		}
	  }()
	  
	  r.Logger.Info().Msg("Waiting to messages")
	  <-forever
}

