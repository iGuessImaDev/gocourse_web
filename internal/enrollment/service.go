package enrollment

import (
	"errors"
	"log"

	"github.com/iGuessImaDev/gocourse_web/internal/course"
	"github.com/iGuessImaDev/gocourse_web/internal/domain"
	"github.com/iGuessImaDev/gocourse_web/internal/user"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log       *log.Logger
		userSrv   user.Service
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(l *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:       l,
		userSrv:   userSrv,
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {
	enroll := domain.Enrollment{
		UserId:   userID,
		CourseId: courseID,
		Status:   "P",
	}

	if _, err := s.userSrv.Get(enroll.UserId); err != nil {
		return nil, errors.New("user id doesn't exist!")
	}

	if _, err := s.courseSrv.Get(enroll.CourseId); err != nil {
		return nil, errors.New("course id doesn't exist!")
	}

	if err := s.repo.Create(&enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return &enroll, nil
}
