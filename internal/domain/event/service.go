package event

type Service interface {
	FindAll() ([]Poster, error)
	FindOne(id int64) (*Poster, error)
	Add(title string, description string, year int64) (*Poster, error)
	Update(id int64, title string, description string, year int64) (*Poster, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Poster, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Poster, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Add(title string, description string, year int64) (*Poster, error) {
	return (*s.repo).Add(title, description, year)
}

func (s *service) Update(id int64, title string, description string, year int64) (*Poster, error) {
	return (*s.repo).Update(id, title, description, year)
}
