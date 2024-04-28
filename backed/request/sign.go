package request

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Register_Msg struct {
	name           string `json:"name" gorm:"name"`
	email          string `json:"email" gorm:"email"`
	password       string `json:"password" gorm:"password"`
	again_password string `json:"again_password" gorm:"again_password"`
}

func Register(c *gin.Context) {
	var data *Register_Msg
	c.BindJSON(&data)
	if len(data.name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}
	if len(data.password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	//err := Rdb.Get
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newuer := Account{
		Name:     data.name,
		Password: string(hashedPassword),
		Email:    data.email,
	}
	db.Create(newuer)
	newMsg := Message{
		Name:             data.name,
		SelfIntroduction: "",
	}
	db.Create(newMsg)
	c.JSON(http.StatusOK, gin.H{})
	c.Redirect(http.StatusOK, "/Use/"+data.name)
}
