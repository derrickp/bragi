package handlers

import (
	"plotsky.dev/bragi/projections"
	"plotsky.dev/bragi/stores"
)

func BuildFetchArtistCounts(store stores.Store) Handler {
	return Handler{
		store,
	}
}

func (handler Handler) FetchArtistCounts(user string) projections.ArtistCounter {
	stored_events, _ := handler.store.GetEvents(user)
	return projections.BuildArtistCounter(stored_events)
}
