package user

import "log"

type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		Create(firstName, lastName, email, phone string) (*User, error)
		GetAll(filter Filters, offset, limit int) ([]User, error)
		Get(id string) (*User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) (*User, error) {
	s.log.Println("create user service")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]User, error) {
	s.log.Println("getall user service")
	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return users, nil
}

func (s service) Get(id string) (*User, error) {
	s.log.Println("get user service")
	user, err := s.repo.Get(id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return user, nil
}

func (s service) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.log.Println(err)
		return err
	}
	return nil
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(id, firstName, lastName, email, phone)
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
