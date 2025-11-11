package validator

type BookPostRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"required,number"`
}

type BookUpdateRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"required,number"`
}
