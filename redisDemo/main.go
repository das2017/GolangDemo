package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/golangDemo/redisDemo/models"
)

func main() {
	fmt.Println("-----------------------welcome to redisdemo-----------------------")
	HashDemo()
}

func HashDemo() {
	fmt.Println("-----------------------welcome to HashDemo-----------------------")
	redisClient := GetRedisClient()
	if redisClient == nil {
		fmt.Errorf("HashDemo redisClient is nil")
		return
	}
	article := models.Article{18, "滴滴出行悄然提价 中国打车市场补贴战或将终结", "", 10, 0}
	articleKey := "150101:article:18"

	redisClient.HMSet(articleKey, models.ToStringDictionary(&article))
	mapOut := redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

	redisClient.HSet(articleKey, "Content", "测试文章内容")
	mapOut = redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

	view, err := redisClient.HIncrBy(articleKey, "Views", 1).Result()
	if err != nil {
		fmt.Printf("\n HIncrBy error=%s ", err)
	} else {
		fmt.Printf("\n HIncrBy Views=%d ", view)
	}
	fmt.Print("\n")

	mapOut = redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

}

func GetRedisClient() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "139.198.13.12:4125",
		Password: "uuid845tylabc123", // no password set
		DB:       0,                  // use default DB
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}
