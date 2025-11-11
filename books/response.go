package books

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type BooksApiResponse struct {
	Status string         `json:"status"`
	Data   []BookResponse `json:"data"`
}

type BookApiResponse struct {
	Status string       `json:"status"`
	Data   BookResponse `json:"data"`
}

type BookDeleteApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
