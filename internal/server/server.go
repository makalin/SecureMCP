package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	router *gin.Engine
}

// NewServer creates a new Server instance
func NewServer() *Server {
	router := gin.Default()
	s := &Server{
		router: router,
	}
	s.setupRoutes()
	return s
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.router.Run(":8080")
}

func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// API routes
	api := s.router.Group("/api/v1")
	{
		api.POST("/scan", s.handleScan)
		api.GET("/reports", s.handleGetReports)
	}
}

func (s *Server) handleScan(c *gin.Context) {
	var request struct {
		Target string `json:"target" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement scan logic
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Scanning target: %s", request.Target),
	})
}

func (s *Server) handleGetReports(c *gin.Context) {
	// TODO: Implement reports retrieval
	c.JSON(http.StatusOK, gin.H{
		"reports": []string{},
	})
}
