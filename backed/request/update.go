package request

import "github.com/gin-gonic/gin"

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
