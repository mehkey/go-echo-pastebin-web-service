package handler

import (
	"github.com/mehkey/restful-go-api/database/internal/datasource"
)

func NewHandler(db datasource.DB) *Handler {
	h := Handler{db}
	return &h
}
