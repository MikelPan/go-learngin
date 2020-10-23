package v1

import (
	"github.com/MikelPan/go-learning/pkg/e"
	"github.com/gin-gonic/gin"
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

