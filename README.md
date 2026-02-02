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

## Learning Path

**Phase 0**: C# Fundamentals — Console apps, syntax, collections  
**Phase 1**: ASP.NET Core Basics — First APIs, routing, DI  
**Phase 2**: Data Access — EF Core, SQL, persistence  
**Phase 3**: API Design — REST, versioning, security  
**Phase 4**: Testing — Unit tests, integration tests  
**Phase 5**: Architecture — Clean architecture, patterns  
**Phase 6**: Performance — Caching, async, real-time  
**Phase 7**: DevOps — Docker, CI/CD  
**Beyond**: Advanced topics, continuous learning

*Assumes programming experience but no C#/backend knowledge.*

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
