package utils

import (
	"breeze-api/api/service"
	"breeze-api/internal/jwt"
	"errors"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 正则验证
// expr 正则表达式
// content 要验证的内容
func CheckRegex(expr, content string) bool {

	r, err := regexp.Compile(expr)
	if err != nil {
		return false
	}

	return r.MatchString(content)
}

// 获取配置信息
func GetConfig(name string) string {

	if config := (&service.Config{}).DetailByName(name); config.Id > 0 {
		return config.Value
	}

	return ""
}

// 获取授权信息
func GetTokenPayload(ctx *fiber.Ctx) (int, error) {
	
	authorization := ctx.Get("Authorization")
	if authorization == "" {
		return 0, errors.New("未授权")
	}

	token := authorization[len("Bearer "):]
	if token == "" {
		return 0, errors.New("未授权")
	}

	payload := jwt.Parse(token)
	if payload == nil || payload.Id <= 0 || time.Now().After(payload.Expire) {
		return 0, errors.New("授权过期")
	}

	return payload.Id, nil
}
