
.PHONY: compose compose-build

compose-build:
	docker compose up -d --build --remove-orphans

compose:
	docker compose up -d
