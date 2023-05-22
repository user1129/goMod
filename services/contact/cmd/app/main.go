package main

import (
	"fmt"
	repository "user1129/microservices/test/pkg/store/postgres"
)

func main() {
	_, err := repository.ConnPostgres("localhost", "guest", 5432, "root", "test")
	if err != nil {
		fmt.Errorf("db bad connection: %s", err)
	}
	fmt.Println("db ok")
}
