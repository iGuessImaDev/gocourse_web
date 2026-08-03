package course

import (
	"log"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *Course) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (repo *repo) Create(course *Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("user created with id ", course.ID)
	return nil
}
