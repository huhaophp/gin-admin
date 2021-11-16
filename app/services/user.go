package services

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/gin-admin/app/http/request"
	"github.com/huhaophp/gin-admin/app/model"
	"github.com/huhaophp/gin-admin/tools/convert"
	"github.com/huhaophp/gin-admin/tools/response"
)

var UserService = &userService{}

type userService struct{}

// GetUsersList GetUsersList
func (*userService) GetUsersList(ctx *gin.Context, request *request.GetUsersListParam) (error, *response.PageResult) {
	param, err := convert.StructToMap(request, "mapstruct")
	if err != nil {
		return err, nil
	}
	err, users, total := model.GetUsersList(param)
	if err != nil {
		return err, nil
	}
	result := response.PageResult{
		List:  users,
		Page:  request.Page,
		Size:  request.Size,
		Total: total,
	}
	return nil, &result
}
