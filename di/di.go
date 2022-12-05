package di

import (
	"log"
	"tracking-server/application"
	"tracking-server/infrastructure"
	"tracking-server/interfaces"
	"tracking-server/shared"

	"go.uber.org/dig"
)

var Container = dig.New()

/**
 * Register container to depedency graph for each module
 */
func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := application.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := interfaces.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := infrastructure.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}
