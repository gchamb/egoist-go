package structs

import (
	"errors"
)

func (req *AuthRequest) ValidateAuthRequest() error {
	if req.Email == "" {
		return errors.New("invalid email")
	}

	if req.Password == ""{
		return errors.New("password cannot be empty")
	}

	if len(req.Password) < 7 {
		return errors.New("password must be at least 7 characters long")
	}

	return nil
}

func (req *UpdateUserRequest) ValidateUpdateUserReq() error {
	if req.CurrentWeight == nil && req.GoalWeight == nil {
		return errors.New("invalid request")
	}

	if req.CurrentWeight != nil && *req.CurrentWeight  < 70 {
		return errors.New("invalid current weight value")
	}

	if req.GoalWeight != nil && *req.GoalWeight  < 70 {
		return errors.New("invalid goal weight value")
	}

	return nil
}