package dto

import "errors"

type TaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	WithPagination bool `json:"with_pagination"`

	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (p *TaskParams) Validate() error {
	if p.WithPagination {
		if p.Page <= 1 {
			p.Page = 1
		}
		if p.PerPage < 10 {
			p.PerPage = 10
		}
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

func (p *TaskParams) Offset() int {
	return (p.Page - 1) * p.PerPage
}