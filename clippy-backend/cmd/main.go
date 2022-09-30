package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
    getAlbums := host.getAlbums()
    
	router := gin.Default()
    router.GET("/albums", )
    router.Run("localhost:8080")
}


