package dto

import "errors"

type GetRepoInput struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

func (i *GetRepoInput) Validate() error {
	if i.Owner == "" {
		return errors.New("owner is required")
	}
	if i.Repo == "" {
		return errors.New("repo is required")
	}
	return nil
}
