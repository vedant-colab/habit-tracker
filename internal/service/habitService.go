package service

import (
	"habit-tracker/internal/dto"
	"habit-tracker/internal/repository"
)

func CreateHabitService(createHabitRequest *dto.CreateHabitRequest) error {
	habit := repository.Habit{
		Username: createHabitRequest.Username,
		Name:     createHabitRequest.Name,
	}

	if err := repository.CreateHabitRepository(habit); err != nil {
		return err
	}
	return nil

}
