package main

import (
	"github.com/arabian9ts/sweeTest/app/builder"
	"github.com/arabian9ts/sweeTest/infrastructure"
)

func main() {
	root, err := builder.InitializeRootController()
	if err != nil {
		panic("failed to initialize controlelrs")
	}
	infrastructure.Router(root).Run()
}
