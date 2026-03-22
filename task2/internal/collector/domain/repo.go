package domain

import (
	"errors"
	"time"
)

type Repo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stars"`
	Forks       int       `json:"forks"`
	CreatedAt   time.Time `json:"created_at"`
}

func (r *Repo) Validate() error {
	if r.Name == "" {
		return DomainErrors
	}
	return nil
}

var DomainErrors = errors.New("Repository not found")
