package main

import "github.com/gin-gonic/gin"

func index(ctx *gin.Context) {
	ctx.HTML()

}

func main() {
	g := gin.Default()
	g.GET("/", index)
}
