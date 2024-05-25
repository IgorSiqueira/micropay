package main

import (
	"github.com/IgorSiqueira/micropay/api/router"
	"github.com/IgorSiqueira/micropay/pkg/env"
)

func main() {
	env.CheckEnvFile()
	port := env.GetEnvValue("PORT")
	router.Initialize(port)
}
