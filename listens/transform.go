package listens

import "plotsky.dev/bragi/spotify"

func FromSpotifyHistoricalListen(id string, historical_listen spotify.HistoricalListen) Listen {
	return Listen{
		ArtistName: historical_listen.ArtistName,
		EndTime:    historical_listen.EndTime,
		ID:         id,
		MsPlayed:   historical_listen.MsPlayed,
		TrackName:  historical_listen.TrackName,
	}
}
