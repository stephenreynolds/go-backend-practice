# Curriculum

A progressive learning path for C#/ASP.NET backend development.

## Phase 1: Core Fundamentals (Weeks 1-2)

### Week 1: ASP.NET Core Basics
- Day 1: Minimal API — Hello World with route parameters
- Day 2: Dependency Injection — Register and inject services
- Day 3: Middleware — Create custom request/response middleware
- Day 4: Configuration — Options pattern, environment-specific config
- Day 5: Model Binding & Validation — DTOs with DataAnnotations
- Day 6: Error Handling — Global exception handler, ProblemDetails
- Day 7: REST (catch-up or review)
- **Weekly Project**: Build a simple TODO API with in-memory storage

### Week 2: Data Access
- Day 1: EF Core Setup — DbContext, entities, migrations
- Day 2: EF Core CRUD — Repository pattern basics
- Day 3: EF Core Relationships — One-to-many, many-to-many
- Day 4: EF Core Queries — Projections, filtering, pagination
- Day 5: Dapper — Raw SQL with micro-ORM
- Day 6: Transactions — Unit of work, transaction scopes
- Day 7: REST
- **Weekly Project**: Refactor TODO API to use SQLite + EF Core

## Phase 2: API Design & Security (Weeks 3-4)

### Week 3: API Best Practices
- Day 1: API Versioning — URL, header, query string strategies
- Day 2: HATEOAS — Hypermedia links in responses
- Day 3: OpenAPI/Swagger — Documentation, schemas, examples
- Day 4: Rate Limiting — Built-in rate limiter middleware
- Day 5: Response Caching — Output caching, ETags
- Day 6: Content Negotiation — JSON, XML, custom formatters
- Day 7: REST
- **Weekly Project**: Build a paginated, versioned, documented API

### Week 4: Authentication & Authorization
- Day 1: JWT Basics — Generate and validate tokens
- Day 2: ASP.NET Identity — User registration, login
- Day 3: Policy-Based Auth — Claims, requirements, handlers
- Day 4: Role-Based Auth — Roles, role checks
- Day 5: OAuth2/OIDC — External providers (Google, GitHub)
- Day 6: Refresh Tokens — Secure token rotation
- Day 7: REST
- **Weekly Project**: Add auth to TODO API (register, login, protect routes)

## Phase 3: Testing & Quality (Weeks 5-6)

### Week 5: Testing
- Day 1: xUnit Basics — Facts, theories, assertions
- Day 2: Mocking — Moq, NSubstitute
- Day 3: Integration Tests — WebApplicationFactory
- Day 4: Test Containers — Spin up real DBs for tests
- Day 5: Fluent Assertions — Readable test assertions
- Day 6: Code Coverage — Generate and analyze reports
- Day 7: REST
- **Weekly Project**: Add comprehensive tests to TODO API (>80% coverage)

### Week 6: Logging & Observability
- Day 1: Structured Logging — Serilog setup
- Day 2: Log Sinks — File, console, Seq
- Day 3: Correlation IDs — Request tracing
- Day 4: Health Checks — Liveness, readiness probes
- Day 5: Metrics — Prometheus, custom counters
- Day 6: Distributed Tracing — OpenTelemetry basics
- Day 7: REST
- **Weekly Project**: Add observability stack to TODO API

## Phase 4: Architecture & Patterns (Weeks 7-8)

### Week 7: Clean Architecture
- Day 1: Project Structure — Layers, dependencies
- Day 2: Domain Layer — Entities, value objects
- Day 3: Application Layer — Use cases, DTOs, interfaces
- Day 4: Infrastructure Layer — Repositories, external services
- Day 5: CQRS Intro — Separate read/write models
- Day 6: MediatR — Commands, queries, handlers
- Day 7: REST
- **Weekly Project**: Rebuild TODO API with Clean Architecture

### Week 8: Advanced Patterns
- Day 1: Domain Events — Raise and handle events
- Day 2: Outbox Pattern — Reliable event publishing
- Day 3: Specification Pattern — Encapsulate query logic
- Day 4: Result Pattern — Error handling without exceptions
- Day 5: Decorator Pattern — Cross-cutting concerns
- Day 6: Feature Flags — Toggle features at runtime
- Day 7: REST
- **Weekly Project**: Add events and feature flags to TODO API

## Phase 5: Performance & Scale (Weeks 9-10)

### Week 9: Performance
- Day 1: Async/Await Deep Dive — ConfigureAwait, ValueTask
- Day 2: Caching Strategies — Memory cache, distributed cache
- Day 3: Redis — Distributed caching with StackExchange.Redis
- Day 4: Response Compression — Gzip, Brotli
- Day 5: Database Optimization — Indexes, query plans
- Day 6: Benchmarking — BenchmarkDotNet
- Day 7: REST
- **Weekly Project**: Optimize TODO API (caching, async, benchmarks)

### Week 10: Real-time & Background Jobs
- Day 1: SignalR Basics — Hubs, clients
- Day 2: SignalR Groups — Targeted messaging
- Day 3: Background Services — IHostedService
- Day 4: Hangfire — Scheduled and recurring jobs
- Day 5: Message Queues — RabbitMQ basics
- Day 6: Worker Services — Long-running processors
- Day 7: REST
- **Weekly Project**: Add real-time notifications to TODO API

## Phase 6: DevOps & Deployment (Weeks 11-12)

### Week 11: Containers
- Day 1: Dockerfile — Multi-stage builds
- Day 2: Docker Compose — Multi-container apps
- Day 3: Health Checks in Docker — Liveness probes
- Day 4: Environment Variables — Secrets, config
- Day 5: Optimizing Images — Size, layers, security
- Day 6: Local Kubernetes — minikube, kind
- Day 7: REST
- **Weekly Project**: Containerize TODO API with Docker Compose

### Week 12: CI/CD
- Day 1: GitHub Actions — Build and test workflow
- Day 2: Code Quality Gates — Linting, formatting
- Day 3: Automated Releases — Semantic versioning
- Day 4: Database Migrations in CI — Safe deployments
- Day 5: Blue/Green Deployments — Zero-downtime concepts
- Day 6: Infrastructure as Code — Basics with Pulumi/Terraform
- Day 7: REST
- **Weekly Project**: Full CI/CD pipeline for TODO API

---

## Beyond Week 12

After completing the curriculum, continue with:
- gRPC and Protobuf
- GraphQL with HotChocolate
- Microservices patterns
- Event sourcing
- Kubernetes deep dive
- Cloud-native patterns (Azure/AWS)
- Security hardening
- API gateways (YARP)

The daily challenges will continue with more advanced topics and real-world scenarios.
