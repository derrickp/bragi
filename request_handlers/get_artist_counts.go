package request_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"plotsky.dev/bragi/projections"
	"plotsky.dev/bragi/stores"
)

type GetArtistCounts struct {
	Store stores.Store
}

func (request_handler GetArtistCounts) Handle(context *gin.Context) {
	user := context.Param("user")
	events, _ := request_handler.Store.GetEvents(user)
	artist_counter := projections.BuildArtistCounter(events)
	context.JSON(http.StatusOK, artist_counter)
}
