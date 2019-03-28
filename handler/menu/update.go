package menu

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/router/middleware/inject"
	"cladmin/service/menu_service"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	menuService := menu_service.Menu{
		Id:       r.Id,
		ParentId: r.ParentId,
		Name:     r.Name,
		Url:      r.Url,
		Perms:    r.Perms,
		Type:     r.Type,
		Icon:     r.Icon,
		OrderNum: r.OrderNum,
	}
	roleList, errNo := menuService.Edit()
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	for _, v := range roleList {
		inject.Obj.Common.RoleAPI.LoadPolicy(v)
	}
	SendResponse(c, nil, nil)
}
