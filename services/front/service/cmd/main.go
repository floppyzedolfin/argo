package main

import (
	"github.com/floppyzedolfin/argo/pkg/logger"
	"github.com/floppyzedolfin/argo/services/front/service/internal"
)

const frontPort = 8411

func main() {
	f := internal.New()
	logger.Log(logger.Info,"front running...")
	f.Listen(frontPort)
}