# Domains
- Users
- Channels
    - Topics/Threads
    - Messages
    - Collaboration Board
    - Drag and Drop


## Dependencies

- Zap Structured Logger
- GO GRPC
- GO OAuth
- Rust Tonic (GRPC CLIENT)
- Redis (Caching & Content Re-Hydration)
- Multiple SurrealDB instances (Rust["For low latency metrics, and optimizations"] & Go)
- Open Telemetry + Grafana & Prometheus (Metrics, User Experience, Optimization)
- Echo (Labstack) to serve OAUTH API
- Flutter (Client)