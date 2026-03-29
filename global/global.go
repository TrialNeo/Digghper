package global

import (
	"Diggpher/pkg/logger"
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
	Log      = logger.Log
)

var (
	JwtSecret    = []byte("JWT_SECRET")
	JwtIssuer    = "Gopher"
	JwtExpiresAt = time.Hour * 72
)
