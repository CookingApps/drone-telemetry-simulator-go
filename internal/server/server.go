package server

import (
	"github.com/CookingApps/drone-telemetry-simulator-go/internal/drone"
	"github.com/CookingApps/drone-telemetry-simulator-go/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	drone  *drone.Drone
}

func NewServer() *Server {
	// Abuja coordinates as starting point
	d := drone.NewDrone("NG-DRONE-001", 9.0765, 7.3986)
	d.StartSimulation()

	s := &Server{
		router: gin.Default(),
		drone:  d,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	h := handlers.NewHandler(s.drone)

	api := s.router.Group("/api/v1")
	{
		api.GET("/status", h.GetStatus)
		api.POST("/takeoff", h.Takeoff)
		api.POST("/land", h.Land)
		api.POST("/waypoint", h.AddWaypoint)
	}

	// WebSocket endpoint
	s.router.GET("/ws/telemetry", h.WebSocketTelemetry)

	// Simple home page
	s.router.GET("/", func(c *gin.Context) {
		c.String(200, "🚁 Naija Drone Telemetry Simulator is running!\n\nTry:\n- GET /api/v1/status\n- POST /api/v1/takeoff\n- WebSocket: /ws/telemetry")
	})
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
