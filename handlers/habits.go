package handlers

import (
	"TrackZen/models"
	"TrackZen/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var db storage.PostgresStorage

func SetStorage(storage storage.PostgresStorage) {
	db = storage
}

func AddHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit.ID = generateID()
	if err := db.AddHabit(habit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, habit)
}

func GetHabits(c *gin.Context) {
	habits, err := db.GetHabits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

func MarkHabitDone(c *gin.Context) {
	id := c.Param("id")
	if err := db.MarkDone(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "habit marked as done"})
}

func generateID() string {
	return uuid.New().String()
}
