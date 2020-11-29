CHECK=\033[32m✔\033[39m
HR=\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#

clean:
	@rm platform.sqlite
	@go run . db migrate
	@go run . db seed

migrate:
	@go run . db migrate

rollback:
	@go run . db rollback

compile:
	@echo "\n${HR}"
	@echo "Parse graphql schema ..."
	@go-bindata -ignore=\.go -pkg=schema -o=gateway/schema/bindata.go gateway/schema/...
	@echo "Parse graphql schema...             ${CHECK} Done"
	@echo "\n${HR}"

run: compile
	@go run . api
