package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"ormuco.go/internal/util"
	"ormuco.go/pkg/models"
)

func (server *HTTPServer) GetAllCacheLRU(w http.ResponseWriter, r *http.Request) {
	var LRUResponse []models.GetLruResponse
	for key, value := range server.cache.cache {
		LRUResponse = append(LRUResponse, models.GetLruResponse{Value: value.value.(string), Expiration: value.expiration, Key: key})

	}
	keys, err := server.redis.Keys(r.Context(), "*").Result()
	if err != nil {
		fmt.Println("Error getting keys:", err)
		return
	}
	for _, key := range keys {
		value, err := server.redis.Get(r.Context(), key).Result()
		if err != nil {
			util.ResponseWithError(w, 400, "Error key not found")

		}
		var convertRedis models.GetLruResponse
		json.Unmarshal([]byte(value), convertRedis)

		LRUResponse = append(LRUResponse, convertRedis)
	}
	util.RespondWithJSON(w, 201, LRUResponse)
}

func (server *HTTPServer) GetLRU(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	value, ok := server.cache.Get(key)
	if !ok {
		redisData, err := server.redis.Get(r.Context(), key).Result()
		if err != nil {
			util.ResponseWithError(w, 400, "Error key not found")
		}
		util.RespondWithJSON(w, 201, redisData)
		return
	}

	util.RespondWithJSON(w, 201, models.GetLruResponse{
		Value:      value.value.(string),
		Expiration: value.expiration,
		Key:        key,
	})

}

func (server *HTTPServer) SetLRU(w http.ResponseWriter, r *http.Request) {
	var createLruRequest models.CreateLruRequest
	err := json.NewDecoder(r.Body).Decode(&createLruRequest)

	if err != nil {
		util.ResponseWithError(w, 400, "Invalid request payload")
		return
	}

	lru := server.cache.Set(createLruRequest.Key, createLruRequest.Value, server)
	if lru == nil {
		util.ResponseWithError(w, 500, "Error key not found")
		return
	}

	lruResponse := models.CreateLruResponse{
		Value:      lru.value.(string),
		Expiration: lru.expiration,
	}

	err = server.redis.Set(r.Context(), createLruRequest.Key, util.ConvertToRedis(lruResponse), time.Duration(server.config.CacheTimeExpiration)*time.Minute).Err()
	if err != nil {
		util.ResponseWithError(w, 500, "Error setting key")
	}

	util.RespondWithJSON(w, 201, lruResponse)

}

func (server *HTTPServer) GetDocs(w http.ResponseWriter, r *http.Request) {
	dir, err := filepath.Abs(".")
	if err != nil {
		http.Error(w, "Failed to get current directory", http.StatusInternalServerError)
		return
	}
	fmt.Println(dir)
	yamlData, err := ioutil.ReadFile(filepath.Join(dir, "backend/docs/doc.yaml"))
	if err != nil {
		http.Error(w, "Failed to read YAML file", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/yaml")

	// Write the YAML data to the response
	w.Write(yamlData)
}

type GeoCache struct {
	mu       sync.RWMutex
	cache    map[string]*cacheItem
	capacity int
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCache(capacity int) *GeoCache {
	return &GeoCache{
		cache:    make(map[string]*cacheItem),
		capacity: capacity,
	}
}

func (c *GeoCache) Set(key string, value interface{}, server *HTTPServer) *cacheItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if _, ok := c.cache[key]; ok {
		c.cache[key].value = value
		if server.config.CacheExpiration {
			c.cache[key].expiration = time.Now().Add(time.Duration(server.config.CacheTimeExpiration) * time.Minute)

		}
		return c.cache[key]
	}

	if len(c.cache)+1 > c.capacity {
		var lruExpiration time.Time
		var lruKey string

		for k, v := range c.cache {
			if v.expiration.Before(lruExpiration) {
				lruKey = k
				lruExpiration = v.expiration
			}
		}

		delete(c.cache, lruKey)
	}

	c.cache[key] = &cacheItem{
		value: value,
		expiration: func() time.Time {
			if server.config.CacheExpiration {
				return time.Now().Add(time.Duration(server.config.CacheTimeExpiration) * time.Minute)
			}

			return time.Time{}

		}(),
	}

	return c.cache[key]
}

func (c *GeoCache) Get(key string) (*cacheItem, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return item, true
}

func (c *GeoCache) Delete(key string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	delete(c.cache, key)
}

func (c *GeoCache) ClearCacheExpiration() {
	var lruExpiration time.Time
	for k, v := range c.cache {
		if v.expiration.Before(lruExpiration) {
			delete(c.cache, k)
		}
	}

}
