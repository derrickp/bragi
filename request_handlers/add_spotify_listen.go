package request_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"plotsky.dev/bragi/commands"
	"plotsky.dev/bragi/handlers"
	"plotsky.dev/bragi/spotify"
	"plotsky.dev/bragi/stores"
)

type AddSpotifyListen struct {
	Store stores.Store
}

func (request_handler AddSpotifyListen) Handle(context *gin.Context) {
	var json spotify.HistoricalListen
	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := context.Param("user")
	command := commands.Command{
		ID: xid.New().String(),
		Data: commands.AddSpotifyListen{
			UserId:           user_id,
			HistoricalListen: json,
		},
	}

	handler := handlers.BuildAddListen(request_handler.Store)
	handler.AddListen(command)
	context.JSON(http.StatusOK, gin.H{"result": "success"})
}
