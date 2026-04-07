package grpc_client

import (
	"context"
	"fmt"
	"time"

	"github.com/marina-popova11/golang-course/task2/gateway/internal/dto"
	pb "github.com/marina-popova11/golang-course/task2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectorInterface interface {
	GetRepoInfo(ctx context.Context, owner, repo string) (dto.GetRepoOutput, error)
}

type Client struct {
	client pb.CollectorServiceClient
	conn   *grpc.ClientConn
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func NewClient(ctx context.Context, address string) (*Client, error) {
	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	return &Client{
		client: pb.NewCollectorServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *Client) GetRepoInfo(ctx context.Context, owner, repo string) (dto.GetRepoOutput, error) {
	resp, err := c.client.GetRepoInfo(ctx, &pb.GetRepoRequest{
		Owner: owner,
		Repo:  repo,
	})
	if err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("grpc client call: %w", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, resp.CreatedAt)

	return dto.GetRepoOutput{
		Name:        resp.Name,
		Description: resp.Description,
		Stars:       int(resp.Stars),
		Forks:       int(resp.Forks),
		CreatedAt:   createdAt,
	}, nil
}
