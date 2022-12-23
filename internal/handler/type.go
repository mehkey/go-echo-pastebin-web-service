package handler

import (
	"github.com/mehkey/restful-go-api/database/internal/datasource"
)

type Handler struct {
	datasource.DB
}

type Message struct {
	Data string `json:"data"`
}
