package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/marina-popova11/golang-course/task2/gateway/internal/dto"
	"github.com/marina-popova11/golang-course/task2/gateway/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	uc *usecase.RepoUsecase
}

func NewHandler(uc *usecase.RepoUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(router *chi.Mux) {
	router.Get("/api/v1/repos/{owner}/{repo}", h.GetRepoInfo)
}

func (h *Handler) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	input := dto.GetRepoInput{
		Owner: chi.URLParam(r, "owner"),
		Repo:  chi.URLParam(r, "repo"),
	}

	output, err := h.uc.GetRepo(r.Context(), input)
	if err != nil {
		statusCode, errorMsg := mapGRPCToHTTPStatus(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(dto.ErrorResponse{
			Error:   http.StatusText(statusCode),
			Message: errorMsg,
			Code:    statusCode,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func mapGRPCToHTTPStatus(err error) (int, string) {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.NotFound:
			return http.StatusNotFound, st.Message()
		case codes.InvalidArgument:
			return http.StatusBadRequest, st.Message()
		case codes.ResourceExhausted:
			return http.StatusTooManyRequests, st.Message()
		case codes.Unauthenticated:
			return http.StatusUnauthorized, st.Message()
		case codes.PermissionDenied:
			return http.StatusForbidden, st.Message()
		default:
			return http.StatusInternalServerError, st.Message()
		}
	}

	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound, err.Error()
	}
	if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "required") {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusInternalServerError, err.Error()
}
