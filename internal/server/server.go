package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	web *gin.Engine
}

func NewServer() *Server {
	godotenv.Load()
	app := gin.Default()

	// Use CORS middleware
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	return &Server{
		app,
	}
}

func (s *Server) Run() {
	s.MapHandlers(s.web)
	s.web.Run(":8080")
	// s.MapHandlers(s.web)
	// s.web.Listen(fmt.Sprintf("%s:%s", os.Getenv("BIND"), os.Getenv("PORT")))
}
