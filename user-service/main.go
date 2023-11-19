package main

import (
	"fmt"
	"github/Tanmoy037/myntraBackend/user-service/db"
)

func main() {
	fmt.Println("Welcome to user service")

	db.CreateTable()

}
