package main

import (
	"github.com/vietlam/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8002")
}
