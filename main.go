package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"plotsky.dev/bragi/handlers"
	"plotsky.dev/bragi/spotify"
	"plotsky.dev/bragi/stores"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	store := stores.BuildInMemoryStore()

	router.POST("/add_spotify_listen", func(context *gin.Context) {
		var json spotify.Listen
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		handler := handlers.BuildAddListen(&store)
		handler.AddListen(json)
		context.JSON(http.StatusOK, gin.H{"result": "success"})
	})

	router.GET("/events", func(context *gin.Context) {
		stream, _ := store.GetEvents("user-1")
		context.JSON(http.StatusOK, gin.H{"events": stream})
	})

	router.Run()
}
