package auth

import (
	"learn-go-fiber/middlewares"
	"learn-go-fiber/modules/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Server ...
type Server struct {
	db *gorm.DB
	service IService
	handler *handler
}

// NewServer ...
func NewServer(db *gorm.DB) *Server {
	userRepository := user.NewRepository(db)
	service := NewService(userRepository)
	handler := newHandler(service)

	server := &Server{}
	server.db = db
	server.service = service
	server.handler = handler
	return server
}

// MountRoutes ...
func (s *Server) MountRoutes(app fiber.Router) {
	r := app.Group("/auth")
	r.Post("/register", s.handler.RegisterUser)
	r.Post("/login", s.handler.Login)
	r.Get("/profile", middlewares.JwtAuth(s.db), s.handler.Profile)
}