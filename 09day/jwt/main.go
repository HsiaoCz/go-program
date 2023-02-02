package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwt是一种轻量认证模式,服务端认证通过后，会产生一个JSON对象,经过签名后得到一个Token
// 这个Token会发回给用户,用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了

// 1.定义一个用于生成签名的字符串
var mySignKey = []byte("Hello xiaofanyi ")

// GenRegisterCliams 使用默认声明创建JWT

func GenRegisteredCliams() (string, error) {
	// 创建Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //这里设置过期时间
		Issuer:    "HsiaoCz",
	}

	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成签名字符串
	return token.SignedString(mySignKey)
}

// 解析token
func ValidateRegisterCliams(tokenString string) bool {
	//解析token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return mySignKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func main() {

}

// 定制化的Claims
// 我们可以自定义JWT中保存的信息
// 比如我们要存储username信息，我们可以自定义一个MyCliams
// jwt的RegisteredCliams只包含了官方字段
// 需要添加更多信息，我们可以添加到自定义的结构体里面
type CustomCliams struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        //内嵌标准声明
}

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

// 定义一个用于签名的字符串
var CustomSecret = []byte("夏天夏天悄悄过去")

// 生成JWT
// 根据自己的业务封装一个生成JWT的函数
func GenToken(username string) (string, error) {
	// 创建自己的声明
	claims := CustomCliams{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "HsiaoCz", //签发人
		},
	}
	// 使用指定签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// 根据给定的JWT字符串解析数据
func ParseToken(tokenString string) (*CustomCliams, error) {
	//解析token
	//如果是自定义的Cliams结构体则需要使用ParseWithCliams方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomCliams{}, func(t *jwt.Token) (interface{}, error) {
		// 使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}

	//对token对象中的Cliams进行类型断言
	if claims, ok := token.Claims.(*CustomCliams); ok && token.Valid {
		//校验token
		return claims, nil
	}
	return nil, errors.New("invlid token")
}
