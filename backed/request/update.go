package request

import "github.com/gin-gonic/gin"

func update(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	msg := Message{
		Name:             data["name"].(string),
		SelfIntroduction: data["selfIntroduction"].(string),
	}
	newone := Message{}
	db.Model(&Message{}).First(&newone, msg.Name)
	newone = msg
	db.Save(&newone)

}
