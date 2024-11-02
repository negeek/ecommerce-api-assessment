package enums

type UserRole struct{}

const (
	User  = "user"
	Admin = "admin"
)

func (o *UserRole) IsValid(role string) bool {
	switch role {
	case User, Admin:
		return true
	default:
		return false
	}
}
