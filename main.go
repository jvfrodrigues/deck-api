/*
Copyright © 2024 Joao Rodrigues jvfr96@gmail.com
*/
package main

import (
	"github.com/jvfrodrigues/deck-api/cmd"
	_ "github.com/jvfrodrigues/deck-api/docs"
)

//	@title			Deck API
//	@version		1.0
//	@description	Simple REST API that allows the creation and management of card decks
//	@BasePath		/api
func main() {
	cmd.Execute()
}
