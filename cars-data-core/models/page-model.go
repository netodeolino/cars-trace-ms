package models

type Data interface {
	CarModel | []CarModel
}

type Page[T Data] struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Data  T      `json:"data"`
}
