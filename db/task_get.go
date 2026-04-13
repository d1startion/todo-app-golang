package db

func GetTask(id string) (*Task, error) {
	var t Task

	err := db.QueryRow(
		`SELECT id, date, title, comment, repeat
		 FROM scheduler
		 WHERE id = ?`,
		id,
	).Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
