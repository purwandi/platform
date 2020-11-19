CHECK=\033[32m✔\033[39m
HR=\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#

migrate:
	@rel migrate -dir=migrations
run:
	@echo "\n${HR}"
	@echo "Parse graphql schema ..."
	@echo "${HR}\n"
	@go-bindata -ignore=\.go -pkg=schema -o=gateway/schema/bindata.go gateway/schema/...
	@echo "\n${HR}"
	@echo "Parse graphql schema...             ${CHECK} Done"
	@go run . api
