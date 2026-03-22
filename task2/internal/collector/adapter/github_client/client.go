package github_client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v84/github"
	"github.com/marina-popova11/golang-course/task2/internal/collector/domain"
)

type RepoInterface interface {
	GetRepoInfo(ctx context.Context, owner, repo string) (*domain.Repo, error)
}

type Client struct {
	client *github.Client
}

func NewClient(token string) *Client {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	client := github.NewClient(httpClient)

	if token != "" {
		client = client.WithAuthToken(token)
	}

	return &Client{client: client}
}

func (c *Client) GetRepoInfo(ctx context.Context, owner, repo string) (*domain.Repo, error) {
	repository, _, err := c.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("github client get repo: %w", err)
	}

	return &domain.Repo{
		Name:        repository.GetName(),
		Description: repository.GetDescription(),
		Stars:       repository.GetStargazersCount(),
		Forks:       repository.GetForksCount(),
		CreatedAt:   repository.GetCreatedAt().Time,
	}, nil
}
