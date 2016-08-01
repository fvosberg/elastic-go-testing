dev:
	docker-compose up --force-recreate -d --build

# should be called within the dev docker - execute with 'make test-docker'
test:
	go get github.com/fvosberg/elastic-go-testing
	go test github.com/fvosberg/elastic-go-testing

test-docker:
	docker-compose run web test
