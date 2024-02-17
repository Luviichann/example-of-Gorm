package main

import (
	"gormnote/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	routers.DefaultRoutersInit(r)
	r.Run()
}

// github.com/pilu/fresh
