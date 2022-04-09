package myslq

import (
	"gorm.io/gorm"
	"vk-chat-bot/internal/repository"
)

type RepositoryImpl struct {
	db *gorm.DB
}

// NewRepository - функция создания нового MySQL репозитория
func NewRepository(db *gorm.DB) repository.Repository {
	return &RepositoryImpl{db: db}
}
