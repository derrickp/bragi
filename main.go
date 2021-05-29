package main

import (
	"github.com/gin-gonic/gin"
	"plotsky.dev/bragi/request_handlers"
	"plotsky.dev/bragi/stores"
)

func main() {
	store := stores.BuildInMemoryStore()

	router := gin.Default()

	router.GET("/health_check", request_handlers.HealthCheck{}.Handle)
	router.POST("/:user/add_spotify_listen", request_handlers.AddSpotifyListen{Store: store}.Handle)
	router.GET("/:user/events", request_handlers.GetUserEvents{Store: store}.Handle)
	router.GET("/:user/artist_counts", request_handlers.GetArtistCounts{Store: store}.Handle)

	router.Run()
}
