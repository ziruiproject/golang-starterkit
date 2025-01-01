package web

type CartCreateRequest struct {
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartUpdateRequest struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}
