package db

import "errors"

func UpdateDate(date string, id string) error {
	res, err := db.Exec(
		`UPDATE scheduler SET date = ? WHERE id = ?`,
		date, id,
	)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("task not found")
	}

	return nil
}
