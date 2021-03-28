PROJECT=argo
VERSION=1.0.0
GO=GO111MODULE=on go

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
	docker build -t ${PROJECT}-front:${VERSION} build/front
	docker build -t ${PROJECT}-portdomain:${VERSION} build/portdomain

docker-run: docker-build
	docker run -rm ${PROJECT}-front:${VERSION}
	docker run -rm ${PROJECT}-portdomain:${VERSION}
