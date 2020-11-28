package todo

import (
	"learn-go-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Server ...
type Server struct {
	db      *gorm.DB
	service IService
	handler *handler
}

// NewServer ...
func NewServer(db *gorm.DB) *Server {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandler(service)

	server := &Server{}
	server.db = db
	server.service = service
	server.handler = handler
	return server
}

// MountRoutes ...
func (s *Server) MountRoutes(app fiber.Router) {
	r := app.Group("/todos")
	r.Get("/", middlewares.JwtAuth(s.db), s.handler.GetAllTodos)
	r.Get("/:id", middlewares.JwtAuth(s.db), s.handler.GetTodoDetail)
	r.Post("/", middlewares.JwtAuth(s.db), s.handler.CreateNewTodo)
	r.Post("/:id/complete", middlewares.JwtAuth(s.db), s.handler.CompleteTodo)
	r.Delete("/:id", middlewares.JwtAuth(s.db), s.handler.DeleteTodo)
}
