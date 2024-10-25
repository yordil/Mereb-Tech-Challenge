package dtos

import (
	"github.com/yordil/mereb-tech-challenge/internal/domain"
)

type SucessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	User    domain.User   `json:"user"`
}