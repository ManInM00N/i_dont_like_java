package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main/binary"
	"net/http"
)

type Register_Msg struct {
	Name           string `form:"name" json:"name" gorm:"name"`
	Email          string `form:"email" json:"email" gorm:"email"`
	Password       string `form:"password" json:"password" gorm:"password"`
	Again_password string `form:"again_password" json:"again_password" gorm:"again_password"`
	Code           string `form:"code" json:"code" gorm:"code"`
}

func (a *Register_Msg) Msg() string {
	return a.Name + "\n" + a.Email + "\n"
}

func Login(c *gin.Context) {
	data := make(map[string]any)
	c.BindJSON(&data)
	//c.Request.
	if len(data["name"].(string)) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名非法",
		})

		c.Redirect(http.StatusFound, "/login")
		return
	}
	var is Account
	is.Name = data["name"].(string)
	err := db.First(&is).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "账号或密码错误",
		})
		return
	}
	isPassword := bcrypt.CompareHashAndPassword([]byte(is.Password), []byte(data["password"].(string)))
	if isPassword != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "账号或密码错误",
		})
		return
	}
	//c.SetCookie(is.Name)
	c.JSON(http.StatusOK, gin.H{})
	//c.Redirect(http.StatusOK, "/user/"+is.Name)
}
func Register(c *gin.Context) {
	var data *Register_Msg
	//tt := make(map[string]interface{})
	//c.BindJSON(&tt)
	//log.Println(tt)
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	binary.InfoLog.Println(data.Msg())

	if !Rule_name.MatchString(data.Name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名非法",
		})
		c.Redirect(http.StatusFound, "/register")
		return
	}
	if !Rule_password.MatchString(data.Password) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码格式非法",
		})
		return
	}
	if !Rule_email.MatchString(data.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱格式非法",
		})
		c.Redirect(http.StatusFound, "/register")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		c.Redirect(http.StatusFound, "/register")
		return
	}
	if binary.Setting.UseRedis {
		if !binary.IsTrue(data.Email, data.Code) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "验证码错误",
			})
			c.Redirect(http.StatusFound, "/register")
			return
		}
	}

	newuer := Account{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}
	if err := db.First(&newuer).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		//log.Println(errors.Is(err, gorm.ErrRecordNotFound))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "账号已存在",
		})
		c.Redirect(http.StatusFound, "/register")
		return
	}
	db.Create(&newuer)
	newMsg := Message{
		Name: data.Name,
		Url:  data.Name + ".jpg",
	}
	db.Create(&newMsg)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"name":     newMsg.Name,
		"motto":    newMsg.Motto,
		"interest": newMsg.Interest,
		"xueli":    newMsg.Xueli,
		"awards":   newMsg.Awards,
		"groups":   newMsg.Group,
	})
}
