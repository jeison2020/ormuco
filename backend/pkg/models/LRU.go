package models

import "time"

type CreateLruRequest struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type CreateLruResponse struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
}

type GetLruResponse struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
	Key        string    `json:"key"`
}
