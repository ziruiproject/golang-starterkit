package web

type UserCreateRequest struct {
	Name     string `validate:"required,max=70,min=3" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}

type UserUpdateRequest struct {
	Id    int    `json:"id"`
	Name  string `validate:"max=20,min=3" json:"name"`
	Email string `validate:"email" json:"email"`
}
