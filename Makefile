PROJECT=argo
VERSION=1.0.0
GO=go

build:
	${GO} build -o build/front/front.out services/front/cmd
	${GO} build -o build/portdomain/portdomain.out services/portdomain/cmd

test:
	${GO} test ./...

cover:
	${GO} test

check-cover: cover
	${GO} test ./... -cover

docker-build: build
	docker build -t ${PROJECT}-front:${VERSION} build/front
	docker build -t ${PROJECT}-portdomain:${VERSION} build/portdomain

docker-run: docker-build
	docker run -rm ${PROJECT}-front:${VERSION}
	docker run -rm ${PROJECT}-portdomain:${VERSION}
