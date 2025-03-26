package main

import (
	"TrackZen/handlers"
	"TrackZen/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Инициализация PostgreSQL
	store, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer store.Close()

	// 2. Создание таблиц
	if err := store.Init(); err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}

	// 3. Настройка обработчиков
	handlers.SetStorage(*store)

	// 4. Настройка роутера
	r := gin.Default()
	r.POST("/habits", handlers.AddHabit)
	r.GET("/habits", handlers.GetHabits)
	r.PUT("/habits/:id/done", handlers.MarkHabitDone)

	// 5. Запуск сервера
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
