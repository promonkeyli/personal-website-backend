package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_backend.com/m/v2/internal/app/model"
	"web_backend.com/m/v2/utils"
)

type AuthController struct {
}

//	@Summary	用户登录
//	@Produce	json
//	@Param		user	body		model.User	true	"username password login"
//	@Success	200		{object}	model.User	"成功"
//	@Failure	400		{object}	string		"请求错误"
//	@Failure	500		{object}	string		"内部错误"
//	@Router		/login [post]
func (T AuthController) AuthLoginController(c *gin.Context) {
	// 前端录入的用户名密码
	var bodyUser model.User
	if err := c.BindJSON(&bodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB, user := model.QueryUser(bodyUser.UserName)
	fmt.Println(bodyUser)
	if DB.Error != nil {
		HandleError(c, 200, false, DB.Error, time.Now())
	} else {
		p := bodyUser.Password
		if p == user.Password {
			// 签发token
			token, e := utils.GenerateToken(p)
			if e != nil {
				HandleError(c, 200, false, e, time.Now())
			} else {
				HandleOk(c, 200, true, "ok",
					"token："+token, time.Now())
			}

		} else {
			HandleError(c, 200, false, "密码错误，请重新输入！", time.Now())
		}

	}

}
