package api

import (
	"base/gopkg/gins"
	"base/handler/api/demo"
	"base/handler/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

func NewHandler(engine *gin.Engine) gins.Handler {
	return &Handler{
		engine: engine,
	}
}

func (h *Handler) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders)
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))

	g := h.engine.Group("/api", middleware.RequestCapture())
	handlers := []gins.Handler{
		demo.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
