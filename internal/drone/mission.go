package drone

import "github.com/sirupsen/logrus"

type Waypoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

func (d *Drone) AddWaypoint(wp Waypoint) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// For now just log - you can expand to actual path following later
	logrus.Infof("Added waypoint: %.6f, %.6f at altitude %.1fm", wp.Latitude, wp.Longitude, wp.Altitude)
}
