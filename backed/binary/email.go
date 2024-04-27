package binary

import (
	"crypto/tls"
	"fmt"
	. "github.com/ManInM00N/go-tool/statics"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
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
			<h2> 这是您的验证码 %s <h2>
	`
)

func SendOut(dir string) {
	code := rand.Int() % 1000000
	if code < 100000 {
		code += 100000
	}

	newmsg := fmt.Sprintf(message, IntToString(code))
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
