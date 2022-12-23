package handler

import (
	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
)

type Handler struct {
	datasource.DB
}

type Message struct {
	Data string `json:"data"`
}
