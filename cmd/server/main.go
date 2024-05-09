package main

import (
	"fmt"

	"github.com/danzBraham/eniqilo-store/internal/drivers/db"
)

func main() {
	fmt.Println("Hello World")
	pool := db.GetConnectionPool()
	defer pool.Close()
}
