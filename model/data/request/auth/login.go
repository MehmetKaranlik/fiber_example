package auth

import (
	"Fibers-Example/core/helper"
	"encoding/json"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=32"`
}

func (l LoginRequest) ValidatedParse(data []byte) (*LoginRequest, []string) {
	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, []string{err.Error()}
	}
	errList := helper.ValidatorHelper[LoginRequest]{
		Model: l,
	}.Validate()
	if len(errList) > 0 {
		return nil, errList
	}
	return &l, nil
}
