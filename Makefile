
.PHONY: compose-build
compose-build:
	docker compose up -d --build --remove-orphans
