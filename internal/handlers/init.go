package handlers

import "github.com/govnocods/JWT-Authorization/internal/db"

type Handler struct {
	DB *db.DataBase
}

func NewHandler(database *db.DataBase) *Handler {
	return &Handler{
		DB: database,
	}
}
