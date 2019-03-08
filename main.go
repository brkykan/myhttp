package main

import (
	"config"
	"coordinator"
	"flag"
)

func main() {

	parallel := flag.Int("parallel", 10, "Allowed parallel requests limit")
	flag.Parse()
	args := flag.Args()
	limit := *parallel

	configuration := config.Setup(args, limit)
	coordinator.Run(configuration)
}
