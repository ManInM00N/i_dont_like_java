package binary

import (
	"context"
	"crypto/tls"
	"fmt"
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
		<p>Hello</p>
			<p>这封信是由 ManInM00N 发送的。<p>
			<p>您收到这封信，意味着您将使用您的邮箱进行账号注册<p>
			<h2> 这是您的验证码 %s ,有效期为5分钟 <h2>
	`
)

var ctx = context.Background()
var Rdb *redis.Client

func RedisInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
func SendOut(dir string) {
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%06v", rand.Intn(600000))
	newmsg := fmt.Sprintf(message, code)
	err := Rdb.SetEx(ctx, dir, code, time.Minute*5).Err()
	if err != nil {
		log.Fatalln(err)
	}
	m := gomail.NewMessage()
	m.SetBody("text/html", newmsg)
	m.SetHeader("To", dir)
	m.SetHeader("From", m.FormatAddress(userName, "ManInM00N"))
	m.SetHeader("Subject", "账号注册")
	//m.Set
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
func IsTrue(dir string, code string) bool {
	coderow, err := Rdb.Get(ctx, dir).Result()
	if err != nil && err != redis.Nil {
		log.Fatalln(err)
	}
	log.Println(code, coderow)
	if coderow != "" && coderow == code {
		return true
	}
	return false
}
