# Curriculum

An ongoing learning path for C#/ASP.NET backend development.  
**Assumes**: Programming experience, but no C# or backend knowledge.

---

## Phase 0: C# Fundamentals (Week 1-2)

*Get comfortable with C# before touching web APIs.*

### Week 1: C# Basics
- Day 1: Hello World — Console app, `dotnet new`, `dotnet run`
- Day 2: Variables & Types — `int`, `string`, `bool`, `var`, type inference
- Day 3: Control Flow — `if`, `switch`, loops (`for`, `foreach`, `while`)
- Day 4: Methods — Parameters, return types, expression-bodied members
- Day 5: Collections — `List<T>`, `Dictionary<K,V>`, LINQ basics
- **Weekly Project**: CLI tool that manages a list (add, remove, list, search)

### Week 2: C# Intermediate
- Day 1: Classes & Objects — Constructors, properties, methods
- Day 2: Interfaces — Contracts, multiple implementation
- Day 3: Records & Immutability — `record` types, `init` properties
- Day 4: Null Safety — Nullable types, `?.`, `??`, null checks
- Day 5: Async/Await Intro — `Task`, `async`, `await` basics
- **Weekly Project**: Refactor CLI tool with classes, interfaces, async file I/O

---

## Phase 1: ASP.NET Core Basics (Weeks 3-4)

### Week 3: First Web API
- Day 1: Your First API — Minimal API "Hello World"
- Day 2: Route Parameters — `/greet/{name}`, path and query params
- Day 3: HTTP Methods — GET, POST, PUT, DELETE basics
- Day 4: JSON & DTOs — Request/response bodies, serialization
- Day 5: Status Codes — 200, 201, 400, 404, 500 and when to use them
- **Weekly Project**: Simple notes API (in-memory CRUD)

### Week 4: Web API Essentials
- Day 1: Dependency Injection — What it is, why it matters
- Day 2: Services & Interfaces — Injecting your own services
- Day 3: Configuration — appsettings.json, environment variables
- Day 4: Middleware — Request pipeline, custom middleware
- Day 5: Validation — DataAnnotations, manual validation
- **Weekly Project**: Expand notes API with DI, config, validation

---

## Phase 2: Data Access (Weeks 5-6)

### Week 5: Entity Framework Core
- Day 1: EF Core Setup — NuGet, DbContext, first entity
- Day 2: Migrations — Create, apply, rollback
- Day 3: Basic CRUD — Add, find, update, delete
- Day 4: Relationships — One-to-many, navigation properties
- Day 5: Querying — Where, OrderBy, Select, Include
- **Weekly Project**: Notes API with SQLite persistence

### Week 6: Advanced Data Access
- Day 1: Pagination — Skip, Take, total counts
- Day 2: Filtering & Sorting — Dynamic queries
- Day 3: Repository Pattern — Abstracting data access
- Day 4: Transactions — SaveChanges, explicit transactions
- Day 5: Dapper Intro — When and why to use raw SQL
- **Weekly Project**: Add pagination, filtering, repository pattern

---

## Phase 3: API Design (Weeks 7-8)

### Week 7: REST Best Practices
- Day 1: Resource Naming — Nouns, plurals, hierarchies
- Day 2: API Versioning — URL, header, query strategies
- Day 3: Error Responses — ProblemDetails, consistent errors
- Day 4: OpenAPI/Swagger — Auto-generated docs
- Day 5: HATEOAS — Links in responses
- **Weekly Project**: Polish notes API with versioning, docs, error handling

### Week 8: API Security Basics
- Day 1: HTTPS & CORS — Secure transport, cross-origin
- Day 2: Authentication Concepts — Who are you?
- Day 3: JWT Basics — Tokens, claims, validation
- Day 4: Authorization — What can you do?
- Day 5: Protecting Endpoints — [Authorize], policies
- **Weekly Project**: Add JWT auth to notes API

---

## Phase 4: Testing (Weeks 9-10)

### Week 9: Unit Testing
- Day 1: xUnit Basics — Test projects, [Fact], [Theory]
- Day 2: Arrange-Act-Assert — Test structure
- Day 3: Mocking — Moq basics, faking dependencies
- Day 4: Testing Services — Business logic tests
- Day 5: Test Coverage — Measuring what's tested
- **Weekly Project**: Add unit tests to notes API services

### Week 10: Integration Testing
- Day 1: WebApplicationFactory — Spinning up test servers
- Day 2: Testing Endpoints — HTTP client, assertions
- Day 3: Test Databases — In-memory or test containers
- Day 4: Authentication in Tests — Faking auth
- Day 5: CI Integration — Running tests in pipelines
- **Weekly Project**: Add integration tests, 80%+ coverage

---

## Phase 5: Architecture (Weeks 11-12)

### Week 11: Clean Architecture
- Day 1: Layered Architecture — Why layers matter
- Day 2: Domain Layer — Entities, value objects
- Day 3: Application Layer — Use cases, interfaces
- Day 4: Infrastructure Layer — Implementations, external services
- Day 5: Putting It Together — Dependency flow
- **Weekly Project**: Restructure notes API to clean architecture

### Week 12: Patterns & Practices
- Day 1: CQRS Intro — Separate reads and writes
- Day 2: MediatR — Commands, queries, handlers
- Day 3: Result Pattern — Errors without exceptions
- Day 4: Domain Events — Decoupled communication
- Day 5: Feature Flags — Runtime toggles
- **Weekly Project**: Add CQRS/MediatR to notes API

---

## Phase 6: Performance & Scale (Weeks 13-14)

### Week 13: Caching & Performance
- Day 1: Memory Caching — IMemoryCache
- Day 2: Distributed Caching — Redis basics
- Day 3: Response Caching — Output caching
- Day 4: Async Best Practices — Avoiding pitfalls
- Day 5: Profiling — Finding bottlenecks
- **Weekly Project**: Add caching layer to notes API

### Week 14: Background & Real-time
- Day 1: Background Services — IHostedService
- Day 2: Scheduled Jobs — Timers, Hangfire intro
- Day 3: SignalR Basics — Real-time connections
- Day 4: SignalR Hubs — Sending/receiving messages
- Day 5: Combining Techniques — Background + real-time
- **Weekly Project**: Add real-time notifications to notes API

---

## Phase 7: DevOps (Weeks 15-16)

### Week 15: Containers
- Day 1: Docker Basics — Images, containers, Dockerfile
- Day 2: Dockerizing ASP.NET — Multi-stage builds
- Day 3: Docker Compose — Multi-container apps
- Day 4: Environment & Secrets — Config in containers
- Day 5: Health Checks — Container orchestration
- **Weekly Project**: Containerize notes API

### Week 16: CI/CD
- Day 1: GitHub Actions — Workflow basics
- Day 2: Build & Test — Automated pipeline
- Day 3: Code Quality — Linting, formatting
- Day 4: Deployment — Publishing, releases
- Day 5: Database Migrations — Safe CI/CD
- **Weekly Project**: Full CI/CD pipeline

---

## Beyond Phase 7: Continuous Learning

After completing the phases, the curriculum cycles into advanced topics:

- **gRPC & Protobuf** — High-performance RPC
- **GraphQL** — Flexible queries with HotChocolate
- **Microservices** — Service decomposition, communication
- **Event Sourcing** — Append-only event stores
- **Kubernetes** — Container orchestration
- **Cloud Native** — Azure/AWS services
- **Security Hardening** — OWASP, penetration testing
- **API Gateways** — YARP, rate limiting, routing
- **Observability** — OpenTelemetry, distributed tracing
- **Domain-Driven Design** — Deep dive

Daily challenges continue indefinitely with:
- Real-world scenarios and problem-solving
- Code review exercises (spot the bug)
- Refactoring challenges
- System design exercises
- Interview-style problems
- New .NET features as they release

---

## Progression Philosophy

1. **Slow is smooth, smooth is fast** — Master basics before advancing
2. **Build one project throughout** — Notes API evolves with your skills
3. **Repetition is key** — Similar patterns across different contexts
4. **Ship working code** — Done > perfect
5. **Review and iterate** — Learn from feedback
