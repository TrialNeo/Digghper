package global

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	WebApp   *fiber.App
	CONFIG   = new(Config)
	DataBase *gorm.DB
	Redis    *redis.Client
	Log      *zap.Logger
	SugarLog *zap.SugaredLogger
)

var (
	JwtSecret    = []byte("JWT_SECRET")
	JwtIssuer    = "Gopher"
	JwtExpiresAt = time.Hour * 72
)
