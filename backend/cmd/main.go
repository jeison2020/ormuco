package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger/v2"
	"ormuco.go/config"
	_ "ormuco.go/docs"
	"ormuco.go/internal/handler"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	cache := handler.NewCache(config.Capacity)
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
	redis := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDbName,
	})
	go watchForExpiredKeys(redis, cache)
	go webSocketServer()
	server, err := handler.NewHTTPServer(config, chi.NewRouter(), cache, redis, clients)

	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	err = server.Run()

	if err != nil {
		log.Println("error runing server", err)
	}
}

func watchForExpiredKeys(redisClient *redis.Client, cache *handler.GeoCache) {
	pubsub := redisClient.Subscribe(context.Background(), "__keyevent@0__:expired")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Printf("Error receiving message from Redis: %v", err)
			continue
		}

		expiredKey := msg.Payload
		cache.Delete(expiredKey)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(expiredKey))
			if err != nil {
				log.Printf("Error sending WebSocket message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}

	}

}

func webSocketServer() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Always allow connections from all origins
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	clients[conn] = true

	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		log.Printf("Received message from client: %s", message)

		// Send a response back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing WebSocket message:", err)
			break
		}
	}

}
