package web

type UserCreateRequest struct {
	Name     string `validate:"required,min=3,max=70" json:"name"`
	Email    string `validate:"required,email,max=100" json:"email"`
	Password string `validate:"required,min=8,max=100" json:"password"`
}

type UserUpdateRequest struct {
	Id    int    `json:"id"`
	Name  string `validate:"min=3,max=100" json:"name"`
	Email string `validate:"email,max=100" json:"email"`
}
