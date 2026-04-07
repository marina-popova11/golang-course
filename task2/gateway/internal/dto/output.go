package dto

import "time"

type GetRepoOutput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stars"`
	Forks       int       `json:"forks"`
	CreatedAt   time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
