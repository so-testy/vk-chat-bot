package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
	"vk-chat-bot/config"
	chatbotController "vk-chat-bot/endpoints/chat-bot"
	chatbotService "vk-chat-bot/internal/chat-bot"
	"vk-chat-bot/internal/connection/vk"
	myslqRepo "vk-chat-bot/internal/repository/my-slq"
)

func main() {
	// Получаем конфиг настройки
	cfg := config.GetConfig()

	// Подключаемся к БД
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.DB.GetDSN(),
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   true,
		DontSupportForShareClause: true,
	}))
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем репозиторий
	repo := myslqRepo.NewRepository(db)

	// Подключаемся к API VK
	conn := vk.NewConnection()

	// Создаем сервис чат-бота
	chatBotService := chatbotService.NewService(conn, repo)

	// Создаем контроллер
	controller := chatbotController.NewController(chatBotService)

	// Создаем роутер
	router := mux.NewRouter()

	// Добавляем каждому роуто обработку из контроллера
	router.HandleFunc("/api/health", controller.HandleVKCallback)

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.HttpAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
