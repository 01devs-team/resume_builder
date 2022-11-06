.SILENT:
help:
	echo "run 'make front' to serve frontend"
	echo "run 'make back' to serve backend"

front:
	docker compose -f ./frontend/docker-compose.development.yml up

back:
	docker compose -f ./backend/docker-compose.development.yml up
