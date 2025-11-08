package userdto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Kontak   string `json:"kontak" binding:"omitempty,min=5,max=20"`
	RoleID   uint   `json:"role_id" binding:"omitempty"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
	Kontak   string `json:"kontak" binding:"omitempty,min=5,max=15"`
	RoleID   uint   `json:"role_id" binding:"omitempty,min=1"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Kontak   string `json:"kontak"`
	RoleID   uint   `json:"role_id"`
	RoleName string `json:"role_name,omitempty"`
}
