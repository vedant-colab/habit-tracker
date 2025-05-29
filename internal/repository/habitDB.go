package repository

import (
	"fmt"
	database "habit-tracker/internal/db"
)

type Habit struct {
	Username string
	Name     string
	ID       int
}

func CreateHabitRepository(habit Habit) error {
	query := "insert into habits(username, name) values($1, $2);"
	_, err := database.DB.Exec(query, habit.Username, habit.Name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateHabitRepository(habit Habit) error {
	query := `
        UPDATE habits
        SET name = $1
        WHERE username = $2 AND id = $3
    `
	res, err := database.DB.Exec(query, habit.Name, habit.Username, habit.ID)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no habit found with id %d for user %s", habit.ID, habit.Username)
	}

	return nil
}

func DeleteHabitRepository(habit Habit) error {
	query := `DELETE FROM habits WHERE username = $1 and ID = $2`
	res, err := database.DB.Exec(query, habit.Username, habit.ID)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no habit found with id %d for user %s", habit.ID, habit.Username)
	}

	return nil
}
