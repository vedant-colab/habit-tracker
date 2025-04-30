package repository

import database "habit-tracker/internal/db"

type Habit struct {
	Username string
	Name     string
}

func CreateHabitRepository(habit Habit) error {
	query := "insert into habits(username, name) values($1, $2);"
	_, err := database.DB.Exec(query, habit.Username, habit.Name)
	if err != nil {
		return err
	}
	return nil
}
