package web

type DefaultParams struct {
	SortBy  string `json:"sortBy" validate:"required,min=3,max=4,oneof=asc desc" san:"min=3,max=4,trim,xss,upper" in:"query=sort_by"`
	OrderBy string `json:"order_by" validate:"required,min=2,max=20" san:"min=2,max=20,trim,xss" in:"query=order_by"`
	Page    int    `json:"page" validate:"required,min=1,max=100" san:"min=1,max=100,trim,xss" in:"query=page"`
	Limit   int    `json:"limit" validate:"required,min=1,max=1000" san:"min=1,max=1000,trim,xss" in:"query=limit"`
}
