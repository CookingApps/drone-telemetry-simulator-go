package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/CookingApps/drone-telemetry-simulator-go/internal/drone"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Handler struct {
	drone *drone.Drone
}

func NewHandler(d *drone.Drone) *Handler {
	return &Handler{drone: d}
}

func (h *Handler) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, h.drone.GetTelemetry())
}

func (h *Handler) Takeoff(c *gin.Context) {
	success := h.drone.Takeoff()
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Drone takeoff successful", "status": h.drone.GetTelemetry()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Takeoff failed - low battery"})
	}
}

func (h *Handler) Land(c *gin.Context) {
	h.drone.Land()
	c.JSON(http.StatusOK, gin.H{"message": "Drone landing initiated", "status": h.drone.GetTelemetry()})
}

func (h *Handler) WebSocketTelemetry(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		telemetry := h.drone.GetTelemetry()
		if err := conn.WriteJSON(telemetry); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}

func (h *Handler) AddWaypoint(c *gin.Context) {
	var wp drone.Waypoint
	if err := c.ShouldBindJSON(&wp); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.drone.AddWaypoint(wp)
	c.JSON(200, gin.H{"message": "Waypoint added", "waypoint": wp})
}
