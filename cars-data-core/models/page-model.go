package models

type Page[T any] struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Data  T      `json:"data"`
}
