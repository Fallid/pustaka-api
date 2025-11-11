package utils

type BookGetData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type BooksResponse struct {
	Status string        `json:"status"`
	Data   []BookGetData `json:"data"`
}

type BookResponse struct {
	Status string      `json:"status"`
	Data   BookGetData `json:"data"`
}

type BookDeleteApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserRegisterData struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type UserRegisterResponse struct {
	Status string           `json:"status"`
	Data   UserRegisterData `json:"data"`
}

type ErrorResponse struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}
