PROJECT=argo
VERSION=1.0.0
# CGO_ENABLED=0 means we're not looking for C libs when using the network packages (these libs are absent in alpine images)
# GO111MODULE=on means we are using modules and not the old vendor dir and GO*** env variables
GO=GO111MODULE=on CGO_ENABLED=0 go

FRONT=front
PORTDOMAIN=portdomain

local-build:
	${GO} build -o build/front/front.out ./services/front/service/cmd
	${GO} build -o build/portdomain/portdomain.out ./services/portdomain/service/cmd

test:
	${GO} test ./...

cover:
	${GO} test

check-cover: cover
	${GO} test ./... -cover

docker-build: local-build
	docker build -t ${PROJECT}-${FRONT}:${VERSION} build/front
	docker build -t ${PROJECT}-${PORTDOMAIN}:${VERSION} build/portdomain

docker-run: network run-front run-portdomain

network:
	docker network create --subnet=172.18.0.0/16 network-${PROJECT} || true

run-front:
	docker run -d -p 8411:8411 --network network-${PROJECT} --ip 172.18.0.22 -v ${PWD}/testdata:/tmp --rm --name ${FRONT} ${PROJECT}-front:${VERSION}

run-portdomain:
	docker run -d -p 8405:8405 --network network-${PROJECT} --ip 172.18.0.23 --rm --name ${PORTDOMAIN} ${PROJECT}-portdomain:${VERSION}

stop: stop-front stop-portdomain

stop-front:
	docker stop ${FRONT} || true

stop-portdomain:
	docker stop ${PORTDOMAIN} || true
