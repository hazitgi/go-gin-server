package common

type UserCreationInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}
