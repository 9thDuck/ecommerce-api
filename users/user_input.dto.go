package users

type createUserInput struct {
	Username string `json:"username" validate:"required,min=8,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
	Role     int    `json:"role" validate:"required,oneof=1 2 3"`
}

type loginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
