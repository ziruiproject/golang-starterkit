package web

type CartResponse struct {
	User      UserResponse `json:"user_data"`
	CartItems []CartItem   `json:"cart_items"`
}

type CartItem struct {
	ProductResponse
	Quantity int `json:"quantity"`
}
