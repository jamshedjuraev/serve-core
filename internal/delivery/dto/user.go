package dto

import "errors"

type SignupParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *SignupParams) Validate() error {
	if p.Username == "" {
		return errors.New("username is required")
	}

	if p.Password == "" {
		return errors.New("password is required")
	}
	return nil
}