package users

type jwtConfig struct {
}

const (
	ADMIN  role = 0
	SELLER role = 1
	BUYER  role = 2
)

type role int

func (r role) isAdmin() bool {
	return r == ADMIN
}

func (r role) isSeller() bool {
	return r == SELLER
}

func (r role) isUser() bool {
	return r == BUYER
}
