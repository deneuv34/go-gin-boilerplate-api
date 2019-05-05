package validation

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type UserValidation struct {
	User CreateUserRequest `json:"user"`
}
