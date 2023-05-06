# Domains
- Users
- Channels
    - Topics/Threads
    - Messages


## Dependencies

- Zap Structured Logger
- GO GRPC
- Rust Tonic (GRPC CLIENT)
- Redis (Caching & Content Re-Hydration)
- Multiple SurrealDB instances (Rust["For low latency metrics, and optimizations"] & Go)
- Open Telemetry + Grafana & Prometheus (Metrics, User Experience, Optimization)
- Rust OAuth