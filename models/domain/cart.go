package domain

type Cart struct {
	Id        int `db:"id"`
	UserId    int `db:"user_id"`
	ProductId int `db:"product_id"`
	Quantity  int `db:"quantity"`
}
