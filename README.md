# .NET Backend Practice

Daily and weekly C#/ASP.NET backend development practice to stay sharp and improve employability.

## Structure

```
├── daily/           # Daily challenges (<1 hour each)
│   └── YYYY-MM-DD/  # One folder per day
├── weekly/          # Weekly projects (<1 week each)
│   └── week-NN/     # One folder per week
├── CURRICULUM.md    # Learning progression plan
└── PROGRESS.md      # Tracking completed work
```

## Current Focus Areas

1. **ASP.NET Core Fundamentals** — Controllers, routing, middleware, DI
2. **Data Access** — EF Core, Dapper, raw ADO.NET, migrations
3. **API Design** — REST best practices, versioning, documentation
4. **Authentication/Authorization** — JWT, OAuth, Identity, policies
5. **Testing** — Unit tests, integration tests, mocking
6. **Performance** — Caching, async patterns, profiling
7. **Architecture** — Clean architecture, CQRS, domain-driven design
8. **DevOps** — Docker, CI/CD, health checks, logging
9. **Real-time** — SignalR, WebSockets
10. **Modern C#** — Records, pattern matching, source generators

## Rules

- Daily challenges: **< 1 hour** (often 20-30 min)
- Weekly projects: **< 1 week** (a few hours total)
- All code reviewed and critiqued by Magi
- Focus on learning, not perfection
- Ship working code, iterate later

## Getting Started

```bash
# Clone the repo
git clone git@github.com:stephenreynolds/dotnet-backend-practice.git

# Each challenge has its own folder with instructions
cd daily/YYYY-MM-DD
cat README.md  # Read the challenge
dotnet new webapi -n Solution  # Or whatever template fits
# ... solve it ...
git add . && git commit -m "Complete daily challenge YYYY-MM-DD"
git push
```
