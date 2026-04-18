
#  Drone Ground Control Simulator (Golang)

A realistic **drone telemetry simulator** and lightweight **Ground Control Station (GCS)** backend built in Go.

This project simulates a drone flying in real-time and exposes a clean API + WebSocket for monitoring and controlling it — exactly the kind of backend systems used in autonomous drone operations.

## ✨ Features

- Realistic real-time telemetry simulation (GPS position, altitude, speed, battery, heading)
- Live WebSocket streaming for telemetry updates every 2 seconds
- REST API for drone control (Takeoff, Land, Add Waypoint, Get Status)
- Proper concurrency handling using mutex (thread-safe)
- Structured logging with Logrus
- Clean layered architecture
- Full Docker support + Makefile for easy development

## 🛠️ Tech Stack

- **Go** 1.23
- Gin Web Framework
- Gorilla WebSocket
- Logrus (structured logging)
- Docker

## 🚀 Quick Start

### Run Locally
```bash
git clone https://github.com/CookingApps/drone-telemetry-simulator-go.git
cd drone-telemetry-simulator-go

# Run with Go
go run main.go
```

Server will start at `http://localhost:8080`

### Using Docker
```bash
docker build -t drone-simulator .
docker run -p 8080:8080 drone-simulator
```

## 📡 API Endpoints

| Method | Endpoint                  | Description                          |
|--------|---------------------------|--------------------------------------|
| GET    | `/`                       | Welcome message                      |
| GET    | `/api/v1/status`          | Get current drone telemetry          |
| POST   | `/api/v1/takeoff`         | Command drone to takeoff             |
| POST   | `/api/v1/land`            | Command drone to land                |
| POST   | `/api/v1/waypoint`        | Add a waypoint (JSON body)           |
| GET    | `/ws/telemetry`           | WebSocket - Live telemetry streaming |

### Example Waypoint Request
```json
POST /api/v1/waypoint
{
  "latitude": 9.0820,
  "longitude": 7.4100,
  "altitude": 80
}
```

## 📁 Project Structure
```
internal/
├── drone/      → Core drone logic and simulation
├── handlers/   → HTTP and WebSocket handlers
└── server/     → Gin server setup and routing
```

## 🎯 Why I Built This

- Demonstrates understanding of **real-time systems** (very important for drone ground stations)
- Shows safe **concurrency** and state management
- Practices clean API design and scalable backend patterns

This project proves I can build reliable backend systems that handle live drone data — even though I’m still early in my drone software journey, I’m eager to learn and contribute.

## 🔮 Future Improvements (Planned)
- Support for multiple drones (fleet mode)
- Persistent storage with PostgreSQL or Redis
- Basic mission planner with path following
- Integration with PX4/MAVLink simulation
- Authentication & API keys

---

Its CookingApps

Passionate about drones, autonomous systems, and building impactful technology for Africa.

Feel free to star ⭐ the repo if you find it useful!

