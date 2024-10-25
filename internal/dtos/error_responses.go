package dtos

import ("github.com/yordil/mereb-tech-challenge/internal/domain")

type AllUserResponse struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Users    []domain.User `json:"user"`
}


type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}