package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"plotsky.dev/bragi/commands"
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

	router.POST("/:user/add_spotify_listen", func(context *gin.Context) {
		var json spotify.HistoricalListen
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := context.Param("user")

		command := commands.Command{
			ID: xid.New().String(),
			Data: commands.AddSpotifyListen{
				UserId:           user,
				HistoricalListen: json,
			},
		}

		handler := handlers.BuildAddListen(&store)
		handler.AddListen(command)
		context.JSON(http.StatusOK, gin.H{"result": "success"})
	})

	router.GET("/:user/events", func(context *gin.Context) {
		user := context.Param("user")
		stream, _ := store.GetEvents(user)
		context.JSON(http.StatusOK, gin.H{"events": stream})
	})

	router.GET("/:user/artist_counts", func(context *gin.Context) {
		user := context.Param("user")
		handler := handlers.BuildFetchArtistCounts(&store)

		counter := handler.FetchArtistCounts(user)
		context.JSON(http.StatusOK, counter)
	})

	router.Run()
}
