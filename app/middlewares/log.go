package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n\n",
	}

	e.Use(middleware.LoggerWithConfig(logger))
}
