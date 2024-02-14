package auth

import (
	"Fibers-Example/core/helper"
	"encoding/json"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=32"`
}

func (l LoginRequest) ValidatedParse(data []byte) (*LoginRequest, error) {
	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	err = helper.ValidatorHelper[LoginRequest]{
		Model: l,
	}.Validate()
	if err != nil {
		return nil, err
	}
	return &l, nil
}
