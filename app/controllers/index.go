package controllers

import (
	"fmt"
	"gositev2/app/services"
	"gositev2/task"
	"github.com/gin-gonic/gin"
)

type IndexController struct{
	BaseController
}

func (i *IndexController) Index(c *gin.Context) {
	s := &services.IndexServices{}
	var params services.SearchParams
	err := c.ShouldBindQuery(&params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(params)

	if params.Page <= 0 {
		params.Page = 1
	}

	data := s.Search(params)
	i.Render(c, "index", data)
}

func (i *IndexController) Image(c *gin.Context) {
	new(task.ImageTask).SaveImage()
}
