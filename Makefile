migrate:
	@rel migrate -dir=migrations
run:
	@go run main.go serve
gen:
	@go-bindata -ignore=\.go -pkg=schema -o=gateway/schema/bindata.go gateway/schema/...

