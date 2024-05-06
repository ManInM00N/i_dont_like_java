package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func headchange(c *gin.Context) {
	var msg Message
	msg.Name = c.PostForm("name")
	log.Println(c.Request.PostFormValue("name"), c.PostForm("name"))
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "file upload failed",
		})
		log.Println(err)
		return
	}
	src, _ := file.Open()
	defer src.Close()
	out, err := os.Create("fronted/src/assets/images/" + msg.Name + path.Ext(file.Filename))
	//log.Println(out.Name())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusHTTPVersionNotSupported, gin.H{
			"message": "path not found",
		})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "file load failed",
		})

		return
	}

	db.First(&msg)
	msg.Url = msg.Name + path.Ext(file.Filename)
	db.Save(&msg)
	c.JSON(200, gin.H{
		"url": msg.Url,
	})
}
func update(c *gin.Context) {
	//data := make(map[string]interface{})
	//c.BindJSON(&data)
	var msg Message
	c.BindJSON(&msg)
	log.Println(msg.Name, " change message")
	newone := Message{Name: msg.Name}

	db.First(&newone)
	newone.Group = msg.Group
	newone.Interest = msg.Interest
	newone.Motto = msg.Motto
	newone.Xueli = msg.Xueli
	newone.Awards = msg.Awards
	db.Save(&newone)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
func Query(c *gin.Context) {
	require := c.Query("search")
	ip := c.ClientIP()
	log.Println(ip, require)
	if len(require) > 0 {
		var repost []Feature
		db.Where("name like ?", fmt.Sprintf("%%%s%%", require)).Limit(10).Find(&repost)
		log.Println(repost)
		c.IndentedJSON(200, gin.H{
			"message": "ok",
			"res":     repost,
		})
	}

}
func ToComment(c *gin.Context) {
	mm := make(map[string]any)
	c.BindJSON(&mm)
	var newmsg Comment
	newmsg.Name = mm["name"].(string)
	if len(newmsg.Name) == 0 {
		newmsg.Name = "匿名"
	}
	newmsg.Inner = mm["inner"].(string)
	db.Create(&newmsg)
	c.JSON(200, gin.H{
		"res": newmsg,
	})
	log.Println(newmsg)

}
func GetComments(c *gin.Context) {
	var res []Comment
	db.Find(&res)
	c.IndentedJSON(200, gin.H{
		"message": "ok",
		"res":     res,
	})
	log.Println(res)
}
