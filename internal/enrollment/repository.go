package enrollment

import (
	"log"

	"github.com/iGuessImaDev/gocourse_web/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
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

func (repo *repo) Create(enroll *domain.Enrollment) error {
	if err := repo.db.Create(enroll).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("enrollment created with id ", enroll.ID)
	return nil
}
