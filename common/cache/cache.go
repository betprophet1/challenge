package cache

import "project/common/cache/client"

func GetClient() client.Client {
	return client.GetRedisClient()
}
