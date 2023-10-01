package repository

import (
	"database/sql"

	"github.com/Enzo3322/go-api.git/entity"
)

type CourseMySqlRepository struct {
	Db *sql.DB
}

func (c CourseMySqlRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare(`Insert into courses(id, name, description, status) values(?,?,?,?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
