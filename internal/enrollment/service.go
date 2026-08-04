package enrollment

import (
	"log"

	"github.com/iGuessImaDev/gocourse_web/internal/domain"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {
	enroll := domain.Enrollment{
		UserId:   userID,
		CourseId: courseID,
		Status:   "P",
	}

	if err := s.repo.Create(&enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return &enroll, nil
}
