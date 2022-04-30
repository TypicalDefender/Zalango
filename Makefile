ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/go-microservice"

COVERAGE_MIN=70

GO111MODULE=on
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org

export GO111MODULE
export GOPROXY
export GOPRIVATE
export GOSUMDB

.PHONY: all
all: fmt vet build test

.PHONY: setup

.PHONY: build
build:
	mkdir -p out/
	go build -v -o ${APP_EXECUTABLE}

build-for-deploy:
	mkdir -p go-microservice/configs
	GOOS=linux GOARCH=arm64 go build -v -o go-microservice/go-microservice
	cp -rf ./docs ./go-microservice/
	cp ./configs/application.sample.yml ./go-microservice/configs/application.yml

.PHONY: build
build-linux:
	mkdir -p out/
	GOOS=linux GOARCH=arm64 go build -v -o ${APP_EXECUTABLE}

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint: ## Check style mistake
	@if [[ `golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } | wc -l | tr -d ' '` -ne 0 ]]; then \
          golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }; \
          exit 2; \
    fi;

.PHONY: fmtcheck
fmtcheck:
	@gofmt -l -s $(SOURCE_DIRS) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

.PHONY: copy-config
copy-config:
	cp ./configs/application.sample.yml ./configs/application.yml
	cp ./configs/test.sample.yml ./configs/test.yml

.PHONY: copy-ci-config
copy-ci-config:
	cp ./configs/ci.sample.yml ./configs/test.yml

.PHONY: test
test:
	@ENVIRONMENT=test go test -cover -race $(UNIT_TEST_PACKAGES)

.PHONY: test-ci
test-ci: build copy-ci-config test_db.migrate
	ENVIRONMENT=test go test -race -p=1 -cover -coverprofile=coverage-temp.out $(UNIT_TEST_PACKAGES)
	@cat ./coverage-temp.out | grep -v "mock.go" > ./coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -func=coverage.out | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)"; exit 0}}'

.PHONY: db.create
db.create:
	createdb -p $(DB_PORT) -Opostgres -Eutf8 $(DB_NAME)

.PHONY: db.migrate
db.migrate: build
	$(APP_EXECUTABLE) migrate

.PHONY: db.rollback
db.rollback: build
	$(APP_EXECUTABLE) rollback

.PHONY: db.drop
db.drop:
	dropdb -p $(DB_PORT) --if-exists -Upostgres $(DB_NAME)

.PHONY: test_db.create
test_db.create:
	createdb -p $(TEST_DB_PORT) -Opostgres -Eutf8 $(TEST_DB_NAME)

.PHONY: test_db.migrate
test_db.migrate: build
	ENVIRONMENT=test $(APP_EXECUTABLE) migrate

.PHONY: test_db.rollback
test_db.rollback: build
	ENVIRONMENT=test $(APP_EXECUTABLE) rollback

.PHONY: test_db.drop
test_db.drop:
	dropdb -p $(TEST_DB_PORT) --if-exists -Upostgres $(TEST_DB_NAME)

