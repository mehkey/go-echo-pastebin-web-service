package handler

import (
	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
)

func NewHandler(db datasource.DB) *Handler {
	h := Handler{db}
	return &h
}
