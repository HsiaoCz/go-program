package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	// 获取初始化翻译器
	if err := InitTrans("zh"); err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.POST("/singup", Signup)
	r.Run("0.0.0.0:9090")
}

func Signup(c *gin.Context) {
	su := SignUpParam{}
	if err := c.ShouldBind(&su); err != nil {
		//获取validator.ValidationErrors类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型的错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 对validator.ValidationErrors类型的错误翻译处理
		// 使用我们自定义的方法，去掉结构体名称
		// 这里主要是对错误翻译进行处理
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": su,
	})
}

// 这里面有一个问题，错误信息没有本地化
// 这里注册一个全局的翻译器
var trans ut.Translator

// 初始化翻译器
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 没有json tag时，翻译的错误会带上后端定义的字段名称
		// 这里使用一个jsontag的自定义方法，转化成json的字段
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 注册将翻译信息中的后端结构体字段转化成json字段的方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation, SignUpParam{})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// 搞了json tag之后还是有问题
// 我们的后端结构体还是给显示出来了
// 这里我们想个办法给他去掉
// 自定义一个去掉的方法
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// 之前那一套还是有个小问题
// "re_password": "re_password必须等于Password"
// 错误提示里面re_password比较的是后端结构体里面的字段
// 将哪个字段也给翻译成json，这里我们需要自定义一个方法实现
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(SignUpParam)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
