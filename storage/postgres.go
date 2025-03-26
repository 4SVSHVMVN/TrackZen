package storage

import (
	"TrackZen/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=go_user dbname=trackzen password=sawaqwe123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	_, err := s.db.Exec(`
        CREATE TABLE IF NOT EXISTS habits (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            done BOOLEAN DEFAULT FALSE
        )`)
	return err
}

func (s *PostgresStorage) AddHabit(habit models.Habit) error {
	_, err := s.db.Exec(
		"INSERT INTO habits (id, name, done) VALUES ($1, $2, $3)",
		habit.ID, habit.Name, habit.Done)
	return err
}

func (s *PostgresStorage) GetHabits() ([]models.Habit, error) {
	rows, err := s.db.Query("SELECT id, name, done FROM habits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []models.Habit
	for rows.Next() {
		var h models.Habit
		if err := rows.Scan(&h.ID, &h.Name, &h.Done); err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}
	return habits, nil
}

func (s *PostgresStorage) MarkDone(id string) error {
	_, err := s.db.Exec("UPDATE habits SET done = TRUE WHERE id = $1", id)
	return err
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}
