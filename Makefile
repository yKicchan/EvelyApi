
GOAGEN= vendor/github.com/goadesign/goa/goagen/goagen

all: clean gen fmt build

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf public/js
	@rm -rf public/schema
	@rm -rf public/swagger
	@rm -f  EvelyApi

gen:
	@$(GOAGEN) app     -d EvelyApi/design
	@$(GOAGEN) swagger -d EvelyApi/design -o public
	@$(GOAGEN) schema  -d EvelyApi/design -o public
	@$(GOAGEN) client  -d EvelyApi/design
	@$(GOAGEN) js      -d EvelyApi/design -o public

build:
	@go build .

fmt:
	@go fmt EvelyApi/design
	@go fmt EvelyApi/controllers/api
	@go fmt EvelyApi/controllers/mailer
	@go fmt EvelyApi/controllers/parser
	@go fmt EvelyApi/middleware
	@go fmt EvelyApi/models
	@go fmt EvelyApi/models/collections
	@go fmt EvelyApi/models/documents
