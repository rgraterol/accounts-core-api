package main

import (
	i "github.com/rgraterol/accounts-core-api/infrastructure/initializers"
)

func main() {
	run()
}

func run() {
	i.ConfigInitializer()
	i.LoggerInitializer()
	i.DatabaseInitializer()
	i.ServerInitializer()
}
