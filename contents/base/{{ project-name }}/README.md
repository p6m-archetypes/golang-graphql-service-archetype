# {{ project_title }}

Go GraphQL service using gqlgen.

## API

The p6m standard GraphQL surface, served at `/graphql` on the service port:

- type `{{ PrefixName }}` (`{ id, displayName }`)
- queries `{{ prefix_name | camel_case }}(id)` and `{{ prefix_name | camel_case }}s`
- mutations `create{{ PrefixName }}`, `update{{ PrefixName }}`, `delete{{ PrefixName }}`

Health (`/health/readiness`, `/health/liveness`) and Prometheus `/metrics` answer on the
management port.

The schema lives at `graph/schema.graphqls`; resolvers in `graph/schema.resolvers.go`.

## Environment

The platform injects: `SERVER_PORT`, `MANAGEMENT_PORT`, `LOGGING_STRUCTURED`,
`DB_HOST`/`DB_PORT`/`DB_USERNAME`/`DB_PASSWORD`/`DB_DBNAME` (when persistence is selected),
`OTEL_SERVICE_NAME`, and `OTEL_EXPORTER_OTLP_ENDPOINT` (traces export iff set).

## Getting started

### Generate GraphQL code

```bash
make generate
```

This runs gqlgen and produces `graph/generated.go` + `graph/model/models_gen.go`. The Docker
builds run the same codegen themselves, so a clean render builds and serves its API out of the
box — `make generate` is only needed for local, on-host builds.

### Build

```bash
make build
```

### Test

```bash
make test
```

### Run

```bash
make run
```

GraphQL on port `{{ service_port }}`, HTTP management on `{{ management_port }}`.
