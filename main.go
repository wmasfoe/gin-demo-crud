package main

import (
	"gin-demo/router"
)

func main() {
	r := router.Router()

	r.Run(":9999")
}
