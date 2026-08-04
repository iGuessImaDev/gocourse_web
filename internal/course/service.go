package course

import (
	"log"
	"time"

	"github.com/iGuessImaDev/gocourse_web/internal/domain"
)

type (
	Service interface {
		Create(name, startDate, endDate string) (*domain.Course, error)
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Update(id string, name *string, startDate *string, endDate *string) error
		Delete(id string) error
		Count(filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}

	Filters struct {
		Name      string
		StartDate string
		EndDate   string
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*domain.Course, error) {
	startDateParsed, err := time.Parse("2006-01-01", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-01", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	course := &domain.Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(course); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return course, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	s.log.Println("getall course service")
	courses, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return courses, nil
}

func (s service) Get(id string) (*domain.Course, error) {
	s.log.Println("get course service")
	course, err := s.repo.Get(id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return course, nil
}

func (s service) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.log.Println(err)
		return err
	}
	return nil
}

func (s service) Update(id string, name, startDate, endDate *string) error {
	var startDateParsed, endDateParsed *time.Time

	if startDate != nil {
		date, err := time.Parse("2006-01-01", *startDate)
		if err != nil {
			s.log.Println(err)
			return err
		}
		startDateParsed = &date
	}

	if endDate != nil {
		date, err := time.Parse("2006-01-01", *endDate)
		if err != nil {
			s.log.Println(err)
			return err
		}
		endDateParsed = &date
	}
	return s.repo.Update(id, name, startDateParsed, endDateParsed)
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
