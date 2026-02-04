package auth

import (
	"github.com/gofiber/fiber/v2"
)

const ContextUserIDKey = "userID"

func MiddlewareAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			token        string
			tokenInvalid = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":   9999,
				"errMsg": "Token is invalid",
			})
		)
		authHeader := c.Get("Authorization")
		//基本的判断
		if authHeader == "" {
			return tokenInvalid
		}
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		} else {
			token = authHeader
		}
		//解析
		claims, err := parseToken(token)
		if err != nil {
			return tokenInvalid
		}
		c.Locals(ContextUserIDKey, claims.UserID)
		return c.Next()
	}
}

// GetUserIDFromContext 从上下文中安全获取 userID
func GetUserIDFromContext(c *fiber.Ctx) (uint, bool) {
	userID := c.Locals(ContextUserIDKey)
	if userID == nil {
		return 0, false
	}
	if id, ok := userID.(uint); ok {
		return id, true
	}
	return 0, false
}
