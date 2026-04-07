package dto

import "errors"

type GetRepoInput struct {
	Owner string `json:"owner" validate:"required"`
	Repo  string `json:"repo" validate:"required"`
}

func (i *GetRepoInput) Validate() error {
	if i.Owner == "" {
		return errors.New("Owner is required")
	}
	if i.Repo == "" {
		return errors.New("Repo is required")
	}

	return nil
}
