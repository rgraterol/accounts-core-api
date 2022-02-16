package main

import (
	"github.com/rgraterol/transfers-core-api/infrastructure/init/initializers"
)

func main() {
	run()
}

func run() {
	initializers.ConfigInitializer()
	initializers.LoggerInitializer()
	initializers.DatabaseInitializer()
	initializers.ServerInitializer()
}
