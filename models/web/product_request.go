package web

type ProductCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserID      int    `json:"user_id"`
}

type ProductUpdateRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
