package handler

import (
	"context"
	"strings"
	"time"

	"github.com/marina-popova11/golang-course/task2/internal/collector/dto"
	"github.com/marina-popova11/golang-course/task2/internal/collector/usecase"
	pb "github.com/marina-popova11/golang-course/task2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedCollectorServiceServer
	uc *usecase.RepoUsecase
}

func NewHandler(uc *usecase.RepoUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(server *grpc.Server) {
	pb.RegisterCollectorServiceServer(server, h)
}

// GetRepoInfo godoc
// @Summary      Get repository information from GitHub
// @Description  Get detailed information about a GitHub repository including name, description, stars, forks and creation date
// @Tags         repositories
// @Accept       json
// @Produce      json
// @Param        owner   path      string  true  "Repository owner (username or organization)"
// @Param        repo    path      string  true  "Repository name"
// @Success      200  {object}  dto.GetRepoOutput
// @Failure      400  {object}  dto.ErrorResponse  "Invalid input parameters"
// @Failure      404  {object}  dto.ErrorResponse  "Repository not found"
// @Failure      500  {object}  dto.ErrorResponse  "Internal server error"
// @Router       /api/v1/repos/{owner}/{repo} [get]
func (h *Handler) GetRepoInfo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoResponse, error) {
	input := dto.GetRepoInput{
		Owner: req.Owner,
		Repo:  req.Repo,
	}

	repo, err := h.uc.GetRepo(ctx, input)
	if err != nil {
		return nil, mapErrorToGRPCStatus(err)
	}

	return &pb.GetRepoResponse{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       int32(repo.Stars),
		Forks:       int32(repo.Forks),
		CreatedAt:   repo.CreatedAt.Format(time.RFC3339),
	}, nil
}

func mapErrorToGRPCStatus(err error) error {
	errStr := err.Error()

	if strings.Contains(errStr, "404") || strings.Contains(errStr, "not found") {
		return status.Errorf(codes.NotFound, "repository not found: %v", err)
	}
	if strings.Contains(errStr, "required") || strings.Contains(errStr, "validate") {
		return status.Errorf(codes.InvalidArgument, "invalid input: %v", err)
	}
	if strings.Contains(errStr, "rate limit") {
		return status.Errorf(codes.ResourceExhausted, "rate limited: %v", err)
	}

	return status.Errorf(codes.Internal, "internal error: %v", err)
}
