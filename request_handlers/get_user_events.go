package request_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"plotsky.dev/bragi/stores"
)

type GetUserEvents struct {
	Store stores.Store
}

func (request_handler GetUserEvents) Handle(context *gin.Context) {
	user_id := context.Param("user")
	stream, _ := request_handler.Store.GetEvents(user_id)
	context.JSON(http.StatusOK, gin.H{"events": stream})
}
