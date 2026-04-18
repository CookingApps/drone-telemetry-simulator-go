# 🚁 Naija Drone Ground Control Simulator (Golang)

A realistic **drone telemetry simulator** and lightweight **Ground Control Station (GCS)** backend built in Go.

Designed to demonstrate skills relevant to autonomous drone systems in Nigeria and Africa — real-time data streaming, concurrency, REST + WebSocket APIs, and mission control concepts.

Perfect for showing readiness to contribute to companies building drones like **Terra Industries**, **Briech UAS**, **radahTech**, and **Arone Technologies**.

## Features

- Realistic telemetry simulation (GPS, altitude, speed, battery, heading)
- Real-time WebSocket streaming (live updates every 2s)
- REST API for drone control (Takeoff, Land, Add Waypoint, Status)
- Proper concurrency handling with mutex
- Structured logging with Logrus
- Docker support + Makefile for easy development
- Clean, layered architecture (internal folders)

## Tech Stack

- Go 1.23
- Gin Web Framework
- Gorilla WebSocket
- Logrus for logging
- Docker

## Quick Start

```bash
git clone https://github.com/yourusername/drone-telemetry-simulator-go.git
cd drone-telemetry-simulator-go
make run
Server runs on http://localhost:8080
API Endpoints

GET / → Welcome message
GET /api/v1/status → Current telemetry
POST /api/v1/takeoff
POST /api/v1/land
POST /api/v1/waypoint → Body: {"latitude": 9.08, "longitude": 7.40, "altitude": 80}
GET /ws/telemetry → WebSocket live feed

Docker
Bashmake docker
docker run -p 8080:8080 drone-simulator
Project Architecture
textinternal/
├── drone/      # Core drone logic & simulation
├── handlers/   # HTTP + WebSocket handlers
└── server/     # Gin server setup
Why This Project?

Shows understanding of real-time systems (critical for drone GCS)
Handles concurrency safely (multiple clients + simulation)
Demonstrates clean API design and scalability thinking
Directly relevant to drone autonomy, fleet management, and data pipelines

Future Enhancements (I plan to add):

Persistent storage (PostgreSQL/Redis)
Multiple drone support (fleet mode)
Simple mission planner with path following
Integration with PX4/MAVLink simulation (if possible)

Built from Abuja with passion for Nigeria's growing drone industry.
Star ⭐ if you find it useful!
```
