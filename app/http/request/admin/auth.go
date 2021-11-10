package admin

// AuthLoginRequest admin auth login request.
type (
	AuthLoginRequest struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
)
