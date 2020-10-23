package v1


import (
	"github.com/MikelPan/go-learning/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 查询用户
// @Produce  json
// @Tags 用户
// @Router /api/v1/user/lists [get]
func GetUser(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": "Hello,world!",
	})
}

// @Summary 查询用户信息
// @Produce  json
// @Tags 用户
// @Router /api/v1/user/user_info [get]
func GetUserInfo(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": "Hello,world!",
	})
}
