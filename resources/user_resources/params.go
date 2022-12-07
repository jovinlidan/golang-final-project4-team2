package user_resources

type UserRegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UserTopupBalanceRequest struct {
	Balance int64 `json:"balance" validate:"required"`
}

type UserGenerateTokenParam struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

// Role Enum
type Role string

const (
	RoleAdmin    Role = "admin"
	RoleCustomer Role = "customer"
)
