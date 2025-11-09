package books

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Price    int `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title"` // alias as sub_title
}