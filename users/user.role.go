package users

const (
	ADMIN  Role = 0
	SELLER Role = 1
	BUYER  Role = 2
)

type Role int

func (r Role) IsAdmin() bool {
	return r == ADMIN
}

func (r Role) IsSeller() bool {
	return r == SELLER
}

func (r Role) IsBuyer() bool {
	return r == BUYER
}
