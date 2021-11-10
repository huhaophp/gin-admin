package admin

// AuthLoginRequest admin auth login request.
type (
	AuthLoginRequest struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required" `
		Password string `form:"password" binding:"required"`
	}
)
