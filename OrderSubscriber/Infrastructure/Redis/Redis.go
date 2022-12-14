package Redis

import (
	ApplicationInterfaces "OrderSubscriber/Application/Common/Interfaces"
	DomainEnums "OrderSubscriber/Domain/Enums"
	"context"
	"github.com/go-redis/redis/v8"
)

type SRedis struct {
	Client *redis.Client
}

func NewRedis(host string) ApplicationInterfaces.IRedis {
	return &SRedis{
		Client: redis.NewClient(&redis.Options{
			Addr: host,
		}),
	}
}

func (rRedis *SRedis) Subscribe(ctx context.Context, channelName DomainEnums.RedisQueues) <-chan string {
	var subscribe = rRedis.Client.Subscribe(ctx, string(channelName))
	channel := make(chan string)
	go func() {
		for msg := range subscribe.Channel() {
			channel <- msg.Payload
		}
	}()
	return channel
}

func (rRedis *SRedis) Close() error {
	err := rRedis.Client.Close()
	if err != nil {
		return err
	}
	return nil
}
