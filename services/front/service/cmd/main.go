package main

import (
	"github.com/floppyzedolfin/argo/services/front/service/internal"
)

const frontPort = 8411

func main() {
	f := internal.New()
	f.Listen(frontPort)
}