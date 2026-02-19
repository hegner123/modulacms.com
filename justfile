default:
    @just --list

# Generate Go code from templ files
generate:
    templ generate

# Run go vet on all packages
vet:
    go vet ./...

# Run tests
test:
    go test ./...

# Build the binary
build: generate
    go build -o modulacms.com .

# Run dev server (requires CMS_BASE_URL env var)
dev: generate
    go run .

# Build and start with docker compose
up:
    docker compose up --build -d

# Rebuild and restart the running container
docker-dev:
    docker compose up --build -d --force-recreate

# Stop docker compose
down:
    docker compose down
