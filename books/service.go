package books

type Service interface {
	FindAll() ([]Book, error)
	FindById(Id int) (Book, error)
	Create(book BookPostRequest) (Book, error)
	Update(Id int, Book BookUpdateRequest) (Book, error)
	Delete(Id int) (error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindById(Id int) (Book, error) {
	book, err := s.repository.FindById(Id)
	return book, err
}

func (s *service) Create(bookRequest BookPostRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
	}
	newbook, err := s.repository.Create(book)
	return newbook, err
}

func (s *service) Update(Id int, bookRequest BookUpdateRequest) (Book, error) {
	book, _ := s.repository.FindById(Id)

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating

	updateBook, err := s.repository.Update(book)
	return updateBook, err
}

func (s *service) Delete(Id int) (error){
	book, _ := s.repository.FindById(Id)
	err := s.repository.Delete(book)
	return err
}
