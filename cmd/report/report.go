package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

func main() {
	report()
}

func retrieveMessages(redis *redis.Client, sender, receiver string) {
	// redis.HGetAll("")
}

func report() {

	// connect to redis and get the results and send them
	red := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	r := gin.Default()

	r.GET("/message/list", func(c *gin.Context) {
		// sender := c.Param("sender")
		// receiver := c.Param("receiver")
		fmt.Println(red)
		//	content := retrieveMessages(red, sender, receiver)
		c.JSON(200, "content")
	})

	r.Run()
}
