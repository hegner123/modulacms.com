server := "root@45.79.252.65"
app := "modulacms-site"
repo := "https://github.com/hegner123/modulacms.com.git"

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

# Deploy to production server
deploy:
    ssh {{ server }} ' \
        set -e && \
        cd /root/modulacms.com && \
        git pull && \
        echo "Building new image..." && \
        docker build -t {{ app }}:new . && \
        echo "Starting new container..." && \
        docker run -d --name {{ app }}-new -p 5051:5050 \
            -e CMS_BASE_URL=https://api.modulacms.com \
            -e CMS_API_KEY=$(cat /root/modulacms.com/.env.production 2>/dev/null | grep CMS_API_KEY | cut -d= -f2 || echo "") \
            -e PORT=5050 \
            {{ app }}:new && \
        sleep 2 && \
        echo "Health check..." && \
        if curl -sf -o /dev/null -w "%{http_code}" http://localhost:5051/static/favicon.svg | grep -q 200; then \
            echo "Health check passed" && \
            docker stop {{ app }} 2>/dev/null || true && \
            docker rm {{ app }} 2>/dev/null || true && \
            docker stop {{ app }}-new && \
            docker rm {{ app }}-new && \
            docker tag {{ app }}:new {{ app }}:previous 2>/dev/null || true && \
            docker tag {{ app }}:new {{ app }}:latest && \
            docker run -d --name {{ app }} -p 5050:5050 \
                -e CMS_BASE_URL=https://api.modulacms.com \
                -e CMS_API_KEY=$(cat /root/modulacms.com/.env.production 2>/dev/null | grep CMS_API_KEY | cut -d= -f2 || echo "") \
                -e PORT=5050 \
                --restart unless-stopped \
                {{ app }}:latest && \
            docker image prune -f && \
            echo "Deploy complete"; \
        else \
            echo "Health check failed, rolling back" && \
            docker stop {{ app }}-new 2>/dev/null || true && \
            docker rm {{ app }}-new 2>/dev/null || true && \
            docker rmi {{ app }}:new 2>/dev/null || true && \
            exit 1; \
        fi'

# First-time server setup (clone repo)
setup:
    ssh {{ server }} 'git clone {{ repo }} /root/modulacms.com'

# View production logs
logs:
    ssh {{ server }} 'docker logs -f {{ app }}'

# Check production status
status:
    ssh {{ server }} 'docker ps -a | grep {{ app }} || echo "No containers found"'

# Rollback to previous version
rollback:
    ssh {{ server }} ' \
        docker stop {{ app }} 2>/dev/null || true && \
        docker rm {{ app }} 2>/dev/null || true && \
        docker run -d --name {{ app }} -p 5050:5050 \
            -e CMS_BASE_URL=https://api.modulacms.com \
            -e CMS_API_KEY=$(cat /root/modulacms.com/.env.production 2>/dev/null | grep CMS_API_KEY | cut -d= -f2 || echo "") \
            -e PORT=5050 \
            --restart unless-stopped \
            {{ app }}:previous && \
        echo "Rolled back to previous version"'
