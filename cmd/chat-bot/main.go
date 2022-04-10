package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"vk-chat-bot/config"
	chatbotController "vk-chat-bot/endpoints/chat-bot"
	chatbotService "vk-chat-bot/internal/chat-bot"
	vkConn "vk-chat-bot/internal/connection/vk"
)

func main() {
	// Получаем конфиг настройки
	cfg := config.GetConfig()

	//// Подключаемся к БД
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       cfg.DB.GetDSN(),
	//	DefaultStringSize:         255,
	//	DisableDatetimePrecision:  true,
	//	DontSupportRenameIndex:    false,
	//	DontSupportRenameColumn:   true,
	//	DontSupportForShareClause: true,
	//}))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Инициализируем репозиторий
	//repo := myslqRepo.NewRepository(db)

	// Подключаемся к API VK
	vk := api.NewVK(cfg.ApiVK.ApiKey)
	vkCallback := callback.NewCallback()
	vkCallback.ConfirmationKey = cfg.ApiVK.ConfirmationKey
	vkCallback.Title = "Bot"

	conn := vkConn.NewConnection(vk)

	// Создаем сервис чат-бота
	chatBotService := chatbotService.NewService(conn, nil)

	// Создаем контроллер
	controller := chatbotController.NewController(vkCallback, chatBotService)
	controller.Run()

	// Создаем роутер
	router := mux.NewRouter()

	// Добавляем роут на обработку callback
	router.HandleFunc("/callback", vkCallback.HandleFunc)

	//go func() {
	//	if err := vkCallback.AutoSetting(vk, cfg.ApiVK.CallbackUrl); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.HttpAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server start: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
