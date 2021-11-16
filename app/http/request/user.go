package request

type (
	GetUsersListParam struct {
		Page     int    `mapstruct:"page" form:"page,default=1" binding:"required"`
		Size     int    `mapstruct:"size" form:"size,default=20" binding:"required"`
		Phone    string `mapstruct:"phone" form:"phone"`
		Email    string `mapstruct:"email" form:"email"`
		Username string `mapstruct:"username" form:"username"`
		Keywords string `mapstruct:"keywords" form:"keywords"`
	}
)
