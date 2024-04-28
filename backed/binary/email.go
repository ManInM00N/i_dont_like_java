package binary

import (
	"context"
	"crypto/tls"
	"fmt"
	. "github.com/ManInM00N/go-tool/statics"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"time"
)

const (
	host     = "smtp.qq.com"
	port     = 25
	userName = "3141529546@qq.com"
	password = "vndnorjmcosidfba"
	message  = `
		<p>Hello.</p>
			<p>这封信是由 ManInM00N 发送的。<p>
			<p>您收到这封信，意味着您将使用您的邮箱进行账号注册<p>
			<h2> 这是您的验证码 %s ,有效期为5分钟 <h2>
	`
)

var ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:7237",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func SendOut(dir string) {
	rand.Seed(time.Now().Unix())
	code := IntToString(rand.Int()%900000 + 100000)

	newmsg := fmt.Sprintf(message, code)
	//err := Rdb.SetEx(ctx, dir, code, time.Minute*5).Err()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	m := gomail.NewMessage()
	m.SetBody("text/html", newmsg)
	m.SetHeader("To", dir)
	m.SetHeader("From", userName)
	m.SetHeader("Subject", "账号注册")
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
}
