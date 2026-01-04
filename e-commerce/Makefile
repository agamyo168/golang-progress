include .env

.PHONY: migrate
create-migration:
	@goose -s create $(filter-out $@, $(MAKECMDGOALS)) sql
