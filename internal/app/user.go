package app

type UserRole string

const (
	SUPERADMIN UserRole = "SUPERADMIN"
	ADMIN      UserRole = "ADMIN"
	CASHIER    UserRole = "CASHIER"
	CUSTOMER   UserRole = "CUSTOMER"
)

type User struct {
	FirstName       string   `json:"first_name" validate:"required,max=50"`
	LastName        string   `json:"last_name" validate:"required"`
	Username        string   `json:"username" validate:"required"`
	Password        string   `json:"password" validate:"required"`
	ConfirmPassword string   `json:"confirm_password" validate:"required,eqfield=Password"`
	PhoneNumber     string   `json:"phone_number" validate:"omitempty,numeric"`
	Email           string   `json:"email" validate:"omitempty,email"`
	Role            UserRole `json:"role" validate:"required"`
}
