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

func ServeInit() {
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
		log.Println()
		log.Println(1)
		c.BindJSON(&data)
	})
	Users := R.Group("/user")
	account := Account{}
	row, err := db.Table(account.TableName()).Select("Name").Rows()
	for row.Next() {
		var ss string
		err = row.Scan(&ss)
		if err != nil {
			log.Fatalln("failed to get message", err)
		}
		Users.GET("/"+ss, func(c *gin.Context) {
			//if !reloadtime.Stop() {
			//	<-reloadtime.C
			//}
			//reloadtime.Reset(time.Duration(portmsg.Killtime) * time.Minute)
			var msg Message
			db.Table(account.TableName()).Where("Account = ?", ss).First(&msg)

		})
		//Users.POST("/"+ss, Updatee)
	}
	R.Use(Cors())
}
