package user

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/user_service"
	"cladmin/util"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var (
		r  ListRequest
		ps util.PageSetting
	)
	if err := c.BindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	userService := user_service.User{
		Username: r.UserName,
	}
	info, count, err := userService.GetList(ps)
	if err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
