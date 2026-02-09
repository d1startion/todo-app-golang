package api

import (
	"errors"
	"time"

	"todo-app/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	today := now.Format("20060102")

	if task.Date == "" {
		task.Date = today
		return nil
	}

	if _, err := time.Parse("20060102", task.Date); err != nil {
		return errors.New("Некорректная дата")
	}

	if task.Date < today {
		if task.Repeat == "" {
			task.Date = today
			return nil
		}

		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
		task.Date = next
	}

	return nil
}
