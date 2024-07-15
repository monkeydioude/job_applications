.PHONY: test
test:
	docker-compose up -d api --no-recreate
	docker-compose exec api bash -c "go test ./..."