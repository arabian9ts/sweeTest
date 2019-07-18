package main

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/builder"
)

func main() {
	repo, _ := builder.InitializeUserRepository()
	admin, err := repo.GetAdminById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admin)
}
