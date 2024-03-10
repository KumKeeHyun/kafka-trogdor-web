package main

import (
	"log"

	"github.com/KumKeeHyun/kafka-trogdor-web/ui"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ui.Register(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
