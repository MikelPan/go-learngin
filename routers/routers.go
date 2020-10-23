package routers

import (
	_ "github.com/MikelPan/go-learning/docs"
	//"github.com/MikelPan/go-learning/middleware/jwt"
	"github.com/MikelPan/go-learning/pkg/setting"
	"github.com/MikelPan/go-learning/routers/api"
	"github.com/MikelPan/go-learning/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.RunMode)
	//url := ginSwagger.URL("http://localhost:8080/swagger/swagger.json")
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
    {
    	apiv1.GET("/index", v1.Home)
    	apiv1.GET("/user/lists",v1.GetUser)
		apiv1.GET("/user/user_info",v1.GetUserInfo)
	}
	return r
}
