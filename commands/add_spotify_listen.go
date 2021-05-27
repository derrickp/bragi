package commands

import "plotsky.dev/bragi/spotify"

type AddSpotifyListen struct {
	UserId           string                   `json:"userId" binding:"required"`
	HistoricalListen spotify.HistoricalListen `json:"listen" binding:"required"`
}
