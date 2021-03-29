# argo

## Description

_argo_ is a small project which consists of two Golang microservices that can communicate with each other.

## Installing

In order to retrieve the repo, use
```bash
git clone http://github.com/floppyzedolfin/argo.git
```

## Contents
This repo consists of two microservices, that share some common libraries (`log`, for instance), that have some API dependencies (the client needs to know the API of the backend), and some internal dependencies (I split the code into modules for clarity).

The main structure - the definition of a Port - is generated by protoc.

## Building
This repo comes with a `Makefile` that exposes several build targets, including
```bash
make docker-build
```
This will build the binaries, and wrap them inside docker images.

## Running
This repo comes with a `Makefile` that exposes several build targets, including
```bash
make docker-run
```
This will run the docker images in a local network. Their IP addresses + ports are hardcoded (in the Makefile, and in the client of the backend):
- front: 172.18.0.22:8411
- portdomain: 172.18.0.23:8405

## Playing with it
The docker-run target mounts the `testdata` directory into the `/tmp` directory of the `front` container. You can place your json files in there.

To load a file in the database, run
```bash
curl -X POST -H "content-type:application/json" 172.18.0.22:8411/ports -d '{"input":"/tmp/ports.json"}'
```
To get the contents of the database, run
```bash
curl 172.18.0.22:8411/ports
```

## Limitations, choices of implementation, discussion
I've decided to use `fiber` for the front service. I've used it in a small test project on my spare time (http://github.com/floppyzedolfin/square) which provided some source for copy/paste.

The first version of my code loaded the whole input file, which didn't respect the requirement of memory limitation. But it allowed me to test the rest of the endpoints.

I've written unit tests over some sections of the code, but I'd need to use `mockgen` to mock the backend.
This would mean rewriting the whole `handlers` package, but I think it would bring some simplicity - it would, however, be a great opportunity to not spawn a client from `front` to `portdomain` on each endpoint call.

Further steps would be to have integration tests in the CI, and to use a staged build (see http://github.com/floppyzedolfin/square for an example of how I would have done it). And probably, having a real database, rather than something lying in the memory, would be extremely necessary very fast - at least for data consistency. I'd use mongo, because I've already used mongo.
