package routers

import (
	"github.com/gin-gonic/gin"
	"go-learning/pkg/setting"
	"go-learning/routers/api"
	"go-learning/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
    {
    	apiv1.GET("/index", v1.Home)
	}
	return r
}
