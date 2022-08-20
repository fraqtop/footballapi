ARGS=$(filter-out $@, $(MAKECMDGOALS))

up:
	docker-compose up -d

down:
	docker-compose down

ps:
	docker-compose ps
