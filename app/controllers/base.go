package controllers

import (
	"gositev2/app/models"
	"gositev2/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (b *BaseController) Render(c *gin.Context, name string, data interface{}) {
	var m = gin.H{
		"Content": data,
		"Navs": b.ListNavs(),
		"Title": b.GetMainTitle(),
	}
	c.Set("html", common.HTMLRenderData{
		Name: name,
		Data: &common.HTMLData{
			Status:     http.StatusOK,
			RenderData: m,
		},
	})
}

func (b *BaseController) ListNavs() []*models.Nav {
	return new(models.Nav).ListNavs()
}


func (b *BaseController) GetMainTitle() *models.Title {
	return new(models.Title).GetMainTitle()
}
