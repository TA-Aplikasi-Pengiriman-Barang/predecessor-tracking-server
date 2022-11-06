package di

import (
	"go.uber.org/dig"
	"log"
	"tracking-server/application"
	"tracking-server/infrastructure"
	"tracking-server/interfaces"
	"tracking-server/shared"
)

var Container = dig.New()

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
