package request

import (
	"github.com/gin-gonic/gin"
	"log"
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

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func init() {
	R = gin.Default()

	R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	//Users := r.Group("/user"){
	//}

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
		log.Println(2)
		c.BindJSON(&data)
		log.Println(data)
		c.JSON(200, gin.H{})
		if data["email"] != nil {
			SendOut(data["email"].(string))
		}
	})
	Api.POST("/register", func(c *gin.Context) {
		data := make(map[string]interface{})
		log.Println(c.JSON)
		log.Println(1)
		c.BindJSON(&data)
	})
	R.Use(Cors())
}
