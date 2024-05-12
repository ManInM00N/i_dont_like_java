package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	. "main/binary"
	"net/http"
)

var R *gin.Engine

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func ServeInit() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	R = gin.Default()

	R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	vis := 0
	R.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		vis++
		c.JSON(200, gin.H{
			"num": vis,
		})
	})
	Api := R.Group("/api")
	Api.POST("/register/sendcode", func(c *gin.Context) {
		data := make(map[string]interface{})
		c.BindJSON(&data)
		c.JSON(200, gin.H{})
		if data["email"] != nil {
			SendOut(data["email"].(string))
		}
	})
	Api.POST("/headerchange", headchange)
	Api.POST("/login", Login)
	Api.POST("/update", update)
	Api.POST("/register", Register)
	Api.POST("/comment", ToComment)
	Api.GET("/comment", GetComments)
	Api.GET("/ws", WebSocketHandle)
	R.GET("/HomeTown", Query)
	Users := R.Group("/user")
	Users.GET("/:id", func(c *gin.Context) {
		//tokenString := c.GetHeader("Authorization")
		var msg Message
		id := c.Param("id")
		msg.Name = id
		InfoLog.Println("visit ", id)
		err := db.First(&msg).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			DebugLog.Println(err)
			c.JSON(404, gin.H{
				"code":    404,
				"message": "User not found",
			})
			c.Redirect(404, "/")
			return
		}
		c.JSON(200, gin.H{
			"name":     msg.Name,
			"motto":    msg.Motto,
			"interest": msg.Interest,
			"xueli":    msg.Xueli,
			"awards":   msg.Awards,
			"groups":   msg.Group,
			"url":      msg.Url,
		})
	})
	go Manager.BroadcastSend()
	go Manager.Start()
	go Manager.Quit()
	R.Use(Cors())
}
