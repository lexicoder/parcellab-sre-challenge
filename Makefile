default: build

build:
	cd src/; \
	go mod download; \
	go mod verify; \
	CGO_ENABLED=0 go build -o ../parcellab-sre-challenge main.go; \
	cd ..

build-docker:
	docker build -t parcellab-sre-challenge:latest -f Dockerfile src/

init:
	./scripts/make_check_prerequisites.sh
	echo $$?
	if [ $$? -ne 0 ]; then exit 1; fi

test:
	cd src; \
	go test -v ./...; \
	cd ..

run:
	cd src; \
	go run main.go; \
	cd ..

deploy-local: init
	docker build -t parcellab-sre-challenge:latest -f Dockerfile src/
	helm install parcellab-sre-challenge charts/parcellab-sre-challenge

clean:
	helm delete --wait parcellab-sre-challenge
	docker rmi parcellab-sre-challenge:latest
