package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/response"
)

// ValidatorFunc验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//1.解析请求，支持Json数据、表单请求和URL Query
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用multipart标头，参数请使用JSON格式。")
		return false
	}
	//表单验证
	errs := handler(obj, c)
	//3.判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}

func validate(data interface{}, rule govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rule,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	//开始验证
	return govalidator.New(opts).ValidateStruct()
}