# BudgetPilot

Cost-aware AI Execution Engine
- Backend: Go
- Frontend: Next.js
- Storage: PostgreSQL
- Rate limit / cache: Redis

## Monorepo
- cmd/api        -> Go API entrypoint
- internal/      -> backend modules
- apps/web       -> Next.js dashboard
- migrations/    -> database migrations
- infra/         -> docker/deploy/ci files

## MVP Features
- Complexity scoring
- Budget + token limits
- Cost-aware model routing
- Request logs + dashboard APIs