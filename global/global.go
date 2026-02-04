package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

var (
	WebApp   *fiber.App
	CONFIG   = new(Config)
	DataBase *gorm.DB
	Redis    *redis.Client
)

var (
	JwtSecret    = []byte("JWT_SECRET")
	JwtIssuer    = "Gopher"
	JwtExpiresAt = time.Hour * 72
)
