package v1

import (
	"github.com/gin-gonic/gin"
	"go-learning/pkg/e"
	"net/http"
)

func Home(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": "Hello,world!",
	})
}

//func GetUser(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//
//
//
//}
