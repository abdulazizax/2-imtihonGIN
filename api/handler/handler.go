package handler

import (
	"exam/storage"
	"time"
)

type handler struct {
	storage storage.IStorage
}

type HandlerConfig struct {
	Storage storage.IStorage
}

func New(c *HandlerConfig) *handler {
	return &handler{
		storage: c.Storage,
	}
}

func IsDateValid(dateStr string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	return err == nil
}
