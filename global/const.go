package global

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

const (
	AppName = "Digghper"
)

var (
	FbConfig = fiber.Config{
		AppName:     AppName,
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
		Prefork:     false,
	}
)
