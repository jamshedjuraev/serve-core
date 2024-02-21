package dto

import (
	"errors"
)

type CreateTaskParams struct {
	UserID      string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p *CreateTaskParams) Validate() error {
	if p.UserID == "" {
		return errors.New("user_id cannot be empty")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type GetTaskParams struct {
	TaskID int    `json:"task_id"`
	UserID string `json:"-"`
}

func (p *GetTaskParams) Validate() error {
	if p.TaskID <= 0 {
		return errors.New("task_id cannot be less then 1")
	}
	if p.UserID == "" {
		return errors.New("user_id cannot be empty")
	}
	return nil
}

type GetTasksParams struct {
	UserID         string `json:"-"`
	WithPagination bool   `json:"with_pagination"`
	Page           int    `json:"page"`
	PerPage        int    `json:"per_page"`
}

func (p *GetTasksParams) Validate() error {
	if p.WithPagination {
		if p.Page <= 1 {
			p.Page = 1
		}
		if p.PerPage < 10 {
			p.PerPage = 10
		}
	}

	if p.UserID == "" {
		return errors.New("user_id cannot be empty")
	}
	return nil
}

func (p *GetTasksParams) Offset() int {
	return (p.Page - 1) * p.PerPage
}

type UpdateTaskParams struct {
	TaskID      int    `json:"task_id"`
	UserID      string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func (p *UpdateTaskParams) Validate() error {
	if p.TaskID <= 0 {
		return errors.New("task_id cannot be less then 1")
	}
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

type DeleteTaskParams struct {
	TaskID int `json:"task_id"`
	UserID string `json:"-"`
}

func (p *DeleteTaskParams) Validate() error {
	if p.TaskID <= 0 {
		return errors.New("task_id cannot be less then 1")
	}
	return nil
}
