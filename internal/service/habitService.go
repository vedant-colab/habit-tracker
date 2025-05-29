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

func UpdateHabitService(username string, id int, updateHabitRequset *dto.UpdateHabitRequest) error {
	habit := repository.Habit{
		Username: username,
		ID:       id,
		Name:     updateHabitRequset.Name,
	}

	if err := repository.UpdateHabitRepository(habit); err != nil {
		return err
	}
	return nil
}

func DeleteHabitService(username string, id int, updateHabitRequset *dto.UpdateHabitRequest) error {
	habit := repository.Habit{
		Username: username,
		ID:       id,
		Name:     updateHabitRequset.Name,
	}

	if err := repository.DeleteHabitRepository(habit); err != nil {
		return err
	}
	return nil
}
