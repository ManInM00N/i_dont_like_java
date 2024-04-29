package request

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 定义秘钥
var jwtKey = []byte("*******")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 登录成功之后发放token
func ReleaseToken(user Account) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //token的有效期是七天
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token的有效期
			IssuedAt:  time.Now().Unix(),     //token发放的时间
			Issuer:    "chengqiang",          //作者
			Subject:   "user token",          //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //根据前面自定义的jwt秘钥生成token

	if err != nil {
		//返回生成的错误
		return "", err
	}
	//返回成功生成的字符换
	return tokenString, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// 路由请求中间件，前端必须把token放在请求头上，对服务器进行请求验证token成功后，才能访问后续的请求路由
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 authorization header：获取前端传过来的信息的
		tokenString := c.GetHeader("Authorization")

		log.Print("请求token", tokenString)

		//验证前端传过来的token格式，不为空，开头为Bearer
		if tokenString == "" {
			c.JSON(401, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "权限不足",
					"code": 401,
				},
			})
			return
		}
		//解析token:common/jwt.go
		token, claims, err := ParseToken(tokenString)
		//解析失败||解析后的token无效
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "权限不足",
					"code": 401,
				},
			})
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		var user Account
		//查询数据库
		db.First(&user, userId)

		// 验证用户是否存在
		if user.ID == 0 {
			c.JSON(401, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "权限不足",
					"code": 401,
				},
			})
		}
		c.Set("user", user)
		c.Next()
	}
}
