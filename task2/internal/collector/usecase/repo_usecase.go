package usecase

import (
	"context"
	"fmt"

	"github.com/marina-popova11/golang-course/task2/internal/collector/adapter/github_client"
	"github.com/marina-popova11/golang-course/task2/internal/collector/dto"
)

type RepoUsecase struct {
	githubRepo github_client.RepoInterface
}

func NewRepoUsecase(githubRepo github_client.RepoInterface) *RepoUsecase {
	return &RepoUsecase{githubRepo: githubRepo}
}

func (u *RepoUsecase) GetRepo(ctx context.Context, input dto.GetRepoInput) (dto.GetRepoOutput, error) {
	if err := input.Validate(); err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("Input validate: %w", err)
	}

	repo, err := u.githubRepo.GetRepoInfo(ctx, input.Owner, input.Repo)
	if err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("githubRepo.GetRepoInfo: %w", err)
	}

	if err := repo.Validate(); err != nil {
		return dto.GetRepoOutput{}, fmt.Errorf("Repo validate: %w", err)
	}

	return dto.GetRepoOutput{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       repo.Stars,
		Forks:       repo.Forks,
		CreatedAt:   repo.CreatedAt,
	}, nil
}
