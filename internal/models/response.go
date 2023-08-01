package models

import "mini-project/internal/domain"

type Response struct {
	Data []domain.Booking `json:"data"`
}