package request

type (
	GetUsersListParam struct {
		Page     int    `form:"page,default=1" binding:"required"`
		Size     int    `form:"size,default=20" binding:"required"`
		Phone    string `form:"phone"`
		Email    string `form:"email"`
		Username string `form:"username"`
		Keywords string `form:"keywords"`
	}
)
