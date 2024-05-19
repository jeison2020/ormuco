package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"ormuco.go/config"
	"ormuco.go/internal/middlewares"
)

type HTTPServer struct {
	config  config.Config
	route   *chi.Mux
	cache   *GeoCache
	redis   *redis.Client
	clients map[*websocket.Conn]bool
}

func (server *HTTPServer) Run() error {
	return http.ListenAndServe(server.config.ServerAddress, middlewares.Logger(server.setupRoutes()))
}

func NewHTTPServer(config config.Config, router *chi.Mux, cache *GeoCache, redis *redis.Client, clients map[*websocket.Conn]bool) (*HTTPServer, error) {
	server := &HTTPServer{
		config:  config,
		route:   router,
		cache:   cache,
		redis:   redis,
		clients: clients,
	}

	return server, nil

}

func (server *HTTPServer) setupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/healthz", HandleReadiness)
	router.Get("/LRU/{key}", server.GetLRU)
	router.Get("/LRU", server.GetAllCacheLRU)
	router.Post("/LRU", server.SetLRU)
	router.Get("/docs/doc.yaml", server.GetDocs)
	router.Get("/compare/{v1}/{v2}", server.GetVersion)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(server.config.SwaggerUrl),
	))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	server.route.Use(cors.Handler)
	server.route.Mount("/api/v1", router)
	return server.route

}
