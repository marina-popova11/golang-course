package usecase

import (
	"context"
	"fmt"

	"github.com/marina-popova11/golang-course/task2/gateway/internal/adapter/grpc_client"
	"github.com/marina-popova11/golang-course/task2/gateway/internal/dto"
)

type RepoUsecase struct {
	collector grpc_client.CollectorInterface
}

func NewRepoUsecase(collector grpc_client.CollectorInterface) *RepoUsecase {
	return &RepoUsecase{collector: collector}
}

func (u *RepoUsecase) GetRepo(ctx context.Context, input dto.GetRepoInput) (dto.GetRepoOutput, error) {
	if err := input.Validate(); err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("Input validate: %w", err)
	}

	repo, err := u.collector.GetRepoInfo(ctx, input.Owner, input.Repo)
	if err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("githubRepo.GetRepoInfo: %w", err)
	}

	return repo, nil
}
