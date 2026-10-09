package user

type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleCustomer Role = "CUSTOMER"
)

func (r Role) IsValid() bool {
	return r == RoleAdmin || r == RoleCustomer
}
