package article

import (
	. "cladmin/handler"
	"cladmin/pkg/errno"
	"cladmin/service/article_service"
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
	articleService := article_service.Article{
		CateId: r.CateId,
		Title:  r.Title,
	}
	info, count, errNo := articleService.GetList(ps)
	if errNo != nil {
		SendResponse(c, errNo, nil)
		return
	}
	SendResponse(c, nil, util.PageUtil(count, ps.Page, ps.Limit, info))
}
