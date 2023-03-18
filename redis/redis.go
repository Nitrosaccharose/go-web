package redis

import (
	"github.com/redis/go-redis/v9"
)

var clusterClient *redis.ClusterClient

func ConnectRedis() *redis.ClusterClient {
	nodes := []string{
		"localhost:7000",
		"localhost:7001",
		"localhost:7002",
		"localhost:7003",
		"localhost:7004",
		"localhost:7005",
	}
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: nodes,
	})
	return clusterClient
}

func GetRedis() *redis.ClusterClient {
	return clusterClient
}
