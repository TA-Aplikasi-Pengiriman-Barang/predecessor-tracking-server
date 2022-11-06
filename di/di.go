package di

import (
	"go.uber.org/dig"
	"log"
	"tracking-server/shared"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}
