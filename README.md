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
This will run the docker images. 
The next step will be to talk with the services.


## Limitations, choices of implementation, discussion
The first version of this repo will not include a _real_ database, but rather something stored in memory. It's dirty and ugly, but, hey, if it works under 2hours, it works :)
