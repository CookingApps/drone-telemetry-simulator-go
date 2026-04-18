package drone

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Telemetry struct {
	DroneID   string  `json:"drone_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Speed     float64 `json:"speed_kmh"`
	Battery   float64 `json:"battery_percent"`
	Heading   float64 `json:"heading"`
	Status    string  `json:"status"` // IDLE, FLYING, LANDED, LOW_BATTERY
}

type Drone struct {
	ID        string
	Telemetry Telemetry
	mu        sync.RWMutex
	isFlying  bool
}

func NewDrone(id string, startLat, startLon float64) *Drone {
	return &Drone{
		ID: id,
		Telemetry: Telemetry{
			DroneID:   id,
			Latitude:  startLat,
			Longitude: startLon,
			Altitude:  0,
			Speed:     0,
			Battery:   100.0,
			Heading:   45.0,
			Status:    "IDLE",
		},
	}
}

// StartSimulation runs in background and updates telemetry every 2 seconds
func (d *Drone) StartSimulation() {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			d.mu.Lock()
			if d.isFlying && d.Telemetry.Battery > 5 {
				// Simulate movement around Abuja area
				d.Telemetry.Latitude += (rand.Float64() - 0.5) * 0.0015
				d.Telemetry.Longitude += (rand.Float64() - 0.5) * 0.0015
				d.Telemetry.Altitude = 50 + rand.Float64()*30
				d.Telemetry.Speed = 30 + rand.Float64()*40
				d.Telemetry.Battery -= 0.8
				d.Telemetry.Status = "FLYING"
			} else if d.Telemetry.Battery <= 5 {
				d.isFlying = false
				d.Telemetry.Status = "LOW_BATTERY"
			}
			d.mu.Unlock()
		}
	}()
}

func (d *Drone) Takeoff() bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.Telemetry.Battery > 20 {
		d.isFlying = true
		d.Telemetry.Altitude = 15
		d.Telemetry.Status = "FLYING"
		return true
	}
	return false
}

func (d *Drone) Land() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.isFlying = false
	d.Telemetry.Altitude = 0
	d.Telemetry.Speed = 0
	d.Telemetry.Status = "LANDED"
}

func (d *Drone) GetTelemetry() Telemetry {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.Telemetry
}

func (d *Drone) SetWaypoint(lat, lon float64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// In real system this would queue waypoints
	log.Printf("Waypoint set: %.6f, %.6f", lat, lon)
}
