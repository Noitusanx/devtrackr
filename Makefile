.PHONY: docker-up docker-down docker-logs run test-db clean

# Start all Docker services
docker-up:
	@echo "🐳 Starting Docker services..."
	docker-compose up -d
	@echo "✅ Docker services started"
	@echo "📊 pgAdmin: http://localhost:5050 (admin@devtrackr.com / password123)"
	@echo "🐰 RabbitMQ: http://localhost:15672 (admin / password123)"

# Stop all Docker services
docker-down:
	@echo "🛑 Stopping Docker services..."
	docker-compose down

# View logs
docker-logs:
	docker-compose logs -f

# Test database connection
test-db:
	@echo "🧪 Testing database connection..."
	go run cmd/test-db/main.go

# Run the application
run:
	@echo "🚀 Starting DevTrackr API..."
	go run cmd/api/main.go

# Clean up everything
clean:
	docker-compose down -v
	docker system prune -f

# Install dependencies
install:
	go mod tidy
	go mod download