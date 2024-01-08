# Makefile to simplify the build and run process for the microservices

# Define the services
SERVICES := auth-service user-service profile-service

# Build all services
build:
	@$(foreach service,$(SERVICES),docker build -t $(service) ./$(service);)

# Run all services
up:
	docker-compose up

# Stop all services
down:
	docker-compose down

# Remove containers and networks created by 'up'
clean:
	docker-compose down --remove-orphans

# Rebuild and run all services
rebuild: build up

# Remove images created by 'build'
clean-images:
	@$(foreach service,$(SERVICES),docker rmi $(service);)

# Shortcut to remove everything (containers, networks, and images)
clean-all: clean clean-images

