package v1


import (
	"github.com/MikelPan/go-learning/pkg/app"
	"github.com/MikelPan/go-learning/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary 查询指定用户
// @Produce  json
// @Tags 用户
// @accept json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 3, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	//200 请求
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"msg": id,
	})

	//400 请求
	//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//return

	//401 请求
	//appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	//return


	//500请求
	//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	//return
}

// @Summary 查询用户
// @Produce  json
// @Tags 用户
// @accept json
// @Produce  json
// @Param id path int false "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/ [get]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	//200 请求
	//appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	//	"msg": "Hello,world!",
	//})

	//400 请求
	//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//return

	//401 请求
	appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	return


	//500请求
	//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	//return
}

// @Summary 新增用户
// @Produce  json
// @Tags 用户
// @accept json
// @Produce  json
// @Param id path int false "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/ [post]
func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	//200 请求
	//appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	//	"msg": "Hello,world!",
	//})

	//400 请求
	//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//return

	//401 请求
	appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	return


	//500请求
	//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	//return
}

// @Summary 修改用户
// @Produce  json
// @Tags 用户
// @accept json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/{id} [put]
func EditUser(c *gin.Context) {
	appG := app.Gin{C: c}
	//200 请求
	//appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	//	"msg": "Hello,world!",
	//})

	//400 请求
	//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//return

	//401 请求
	appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	return


	//500请求
	//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	//return
}

// @Summary 删除用户
// @Produce  json
// @Tags 用户
// @accept json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	//200 请求
	//appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	//	"msg": "Hello,world!",
	//})

	//400 请求
	//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//return

	//401 请求
	appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	return


	//500请求
	//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	//return
}



