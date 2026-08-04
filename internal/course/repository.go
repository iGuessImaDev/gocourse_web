package course

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/iGuessImaDev/gocourse_web/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Update(id string, name *string, startDate, endDate *time.Time) error
		Delete(id string) error
		Count(filters Filters) (int, error)
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

func (repo *repo) Create(course *domain.Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("user created with id ", course.ID)
	return nil
}

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course

	tx := repo.db.Model(&c)
	tx = applyFilters(tx, filters)
	tx = tx.Offset(offset).Limit(limit)
	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}

	return c, nil
}

func (repo *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}
	result := repo.db.First(&course)
	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}
	return &course, nil
}

func (repo *repo) Delete(id string) error {
	course := domain.Course{ID: id}

	if err := repo.db.Delete(&course).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("course deleted with id ", course.ID)
	return nil
}

func (repo *repo) Update(id string, name *string, startDate, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if err := repo.db.Model(&domain.Course{}).Where("id", id).Updates(values).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	return nil
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	if filters.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", filters.StartDate)
		if err == nil {
			tx = tx.Where("start_date >= ?", startDate)
		}
	}

	if filters.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", filters.EndDate)
		if err == nil {
			tx = tx.Where("end_date < ?", endDate)
		}
	}
	return tx
}
