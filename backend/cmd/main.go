package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger/v2"
	"log"
	"ormuco.go/config"
	_ "ormuco.go/docs"
	"ormuco.go/internal/handler"
)

func main() {
	scheduler := cron.New(cron.WithSeconds())
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	cache := handler.NewCache(config.Capacity)
	_, err = scheduler.AddFunc("* * * * * *", func() {
		cache.ClearCacheExpiration()
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
	scheduler.Start()
	redis.NewClient(&redis.Options{})
	server, err := handler.NewHTTPServer(config, chi.NewRouter(), cache, redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,  // Redis server address
		Password: config.RedisPassword, // No password
		DB:       config.RedisDbName,
	}))
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	err = server.Run()

	if err != nil {
		log.Println("error runing server", err)
	}
}
