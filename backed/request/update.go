package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func update(c *gin.Context) {
	//data := make(map[string]interface{})
	//c.BindJSON(&data)
	var msg Message
	c.BindJSON(&msg)

	newone := Message{Name: msg.Name}

	db.First(&newone)
	newone.Group = msg.Group
	newone.Interest = msg.Interest
	newone.Motto = msg.Motto
	newone.Xueli = msg.Xueli
	newone.Awards = msg.Awards
	db.Save(&newone)

}
func Query(c *gin.Context) {
	require := c.Query("search")
	mod := c.Query("type")
	ip := c.ClientIP()
	log.Println(ip, mod, require)
	if len(require) > 0 {
		var repost []Food
		db.Where("name like ?", fmt.Sprintf("%%%s%%", require)).Limit(10).Find(&repost)
		c.IndentedJSON(200, gin.H{
			"message": "ok",
			"res":     repost,
		})
	}

}
