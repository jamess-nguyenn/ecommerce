package connection

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"os"
)

type RedisClient struct {
	Client *redis.Client
}

func buildAddress() string {
	address := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	return address
}

func getPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

func BuildRedisClientOptionsAsynq() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     buildAddress(),
		Password: getPassword(),
	}
}

func BuildRedisClientOptions() *redis.Options {
	return &redis.Options{
		Addr:     buildAddress(),
		Password: getPassword(),
	}
}

func openConnectionRedis() *redis.Client {
	client := redis.NewClient(
		BuildRedisClientOptions(),
	)

	return client
}

func pingConnectionRedis(client *redis.Client) error {
	if _, err := client.Ping().Result(); err != nil {
		return fmt.Errorf("failed to ping the redis client: %w", err)
	}

	return nil
}

func closeConnectionRedis(client *redis.Client) error {
	if err := client.Close(); err != nil {
		return fmt.Errorf("failed to close the redis client: %w", err)
	}

	return nil
}

func ConnectRedis() (*RedisClient, error) {
	// open a connection to the redis client
	client := openConnectionRedis()

	// ping the redis client to check the connection
	if err := pingConnectionRedis(client); err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) Close() error {
	// close a connection to the redis client
	if err := closeConnectionRedis(r.Client); err != nil {
		return fmt.Errorf("errors occurred while closing the redis client: %v", err)
	}

	return nil
}
