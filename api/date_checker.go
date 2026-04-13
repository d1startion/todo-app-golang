package api

import (
	"errors"
	"time"

	"todo-app/db"
)

const dateLayout = "20060102"

func checkDate(task *db.Task) error {
	now := time.Now()
	today := now.Format(dateLayout)

	if task.Date == "" {
		task.Date = today
		return nil
	}

	if _, err := time.Parse(dateLayout, task.Date); err != nil {
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
