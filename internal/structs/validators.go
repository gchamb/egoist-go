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