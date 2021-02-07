package routes

import (
	"bytes"
	"fmt"
	"gositev2/app/controllers"
	"gositev2/common"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"runtime/debug"
)

func MyRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				renderError(c, err)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}

func MyRenderHTML() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		data, ok := c.Get("html")
		if !ok {
			return
		}
		htmlData, ok := data.(common.HTMLRenderData)
		if !ok {
			return
		}
		fmt.Println("render html")
		view := Engine.HTMLRender.(*ginview.ViewEngine)
		buf := bytes.NewBuffer(nil)
		err := view.RenderWriter(buf, htmlData.Name, htmlData.Data)
		if err != nil {
			renderError(c, err)
			return
		}
		c.Data(http.StatusOK, "text/html", buf.Bytes())
	}
}

func renderError(c *gin.Context, err interface{})  {
	stack := string(debug.Stack())
	reg := regexp.MustCompile("\n")
	stackSlice := reg.Split(stack, -1)

	status := http.StatusInternalServerError
	errorData := common.HTMLData{}
	switch err.(type) {
	case *common.ErrorData:
		v := err.(*common.ErrorData)
		errorData.Err = v
		status = v.Status
	default:
		errText := "server error"
		if err2,ok := err.(error); ok {
			errText = err2.Error()
		}
		if err2,ok := err.(string); ok {
			errText = err2
		}
		errorData.Err = &common.ErrorData{
			ErrorString: errText,
		}
	}
	if status == http.StatusOK {
		errorData.Message = errorData.Err.ErrorString
	} else {
		errorData.Message = http.StatusText(status)
	}
	errorData.Status = status
	errorData.Err.ErrorStack = stackSlice
	b := new(controllers.BaseController)
	errorData.RenderData = gin.H{
		"Navs": b.ListNavs(),
		"Title": b.GetMainTitle(),
	}

	c.HTML(status, "error", errorData)
}
