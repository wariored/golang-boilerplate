# Define the backend, frontend, build, and test targets
.PHONY: backend-start frontend-start backend-build

# Run the backend server
backend-start:
	cd backend && docker compose up -d

backend-build:
	cd backend && docker compose build

backend-bash:
	cd backend && docker compose exec backend /bin/bash

backend-logs:
	cd backend && docker compose logs backend

backend-attach:
	cd backend && docker compose up backend

backend-stop:
	cd backend && docker compose stop backend

backend-down-all:
	cd backend && docker compose down

backend-stop-all:
	cd backend && docker compose stop

# Run the Delve debugger inside the container
backend-debug:
	cd backend && docker compose exec backend dlv debug --headless --listen=:2345 --log

# Run the frontend server
frontend-start:
	cd frontend && yarn install && yarn start 
