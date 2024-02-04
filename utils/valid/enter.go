package valid

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	trans ut.Translator
)

func init() {
	_ = InitTrans("zh")
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改 gin 框架中的 Validator 引擎属性，实现自定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取 json tag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			fieldName := fld.Name
			// 先尝试获取 label
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			if name == "" {
				// 没有 label 就用 json
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				return ","
			}
			return fmt.Sprintf("%s,%s", fieldName, name)
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
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
		return nil
	}
	return nil
}

// Error 参数校验错误时返回错误信息
func Error(err error) (ret string) {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return err.Error()
	}
	for _, e := range validationErrors {
		msg := e.Translate(trans)
		oldFieldName := e.Field()
		_list := strings.Split(oldFieldName, ",")
		var fieldName string
		if len(_list) > 1 {
			fieldName = _list[0]
		}
		msg = strings.ReplaceAll(msg, fieldName+",", "")
		ret += msg + ";"

	}
	return ret
}

// InValidError 参数校验错误时返回错误信息和数据，数据是校验错误对应的字段
func InValidError(err error, obj any) (ret string, data map[string]string) {
	data = map[string]string{}
	getObj := reflect.TypeOf(obj)
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error(), data
	}

	for _, e := range validationErrors {
		msg := e.Translate(trans)
		oldFieldName := e.Field()
		_list := strings.Split(oldFieldName, ",")
		var fieldName string
		if len(_list) > 1 {
			fieldName = _list[0]
		}
		filed, ok := getObj.Elem().FieldByName(fieldName)
		if ok {
			msg = strings.ReplaceAll(msg, fieldName+",", "")
			jsonLabel, jsonOk := filed.Tag.Lookup("json")
			if jsonOk {
				data[jsonLabel] = msg
			}
		}
		ret += msg + ";"
	}

	return ret, data
}
