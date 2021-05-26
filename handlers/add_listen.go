package handlers

import (
	"github.com/rs/xid"
	"plotsky.dev/bragi/events"
	"plotsky.dev/bragi/listens"
	"plotsky.dev/bragi/projections"
	"plotsky.dev/bragi/spotify"
	"plotsky.dev/bragi/stores"
)

type Handler struct {
	store stores.Store
}

func BuildAddListen(store stores.Store) Handler {
	return Handler{
		store,
	}
}

func (handler Handler) AddListen(listen spotify.Listen) {
	stored_events, version := handler.store.GetEvents("user-1")
	tracker := projections.BuildListenTracker(stored_events)
	id := listen.EndTime + "-" + listen.ArtistName + "-" + listen.TrackName
	_, ok := tracker.GetListen(id)

	if ok {
		return
	}

	event := events.Event{
		ID: xid.New().String(),
		Data: events.ListenAdded{
			Listen: listens.Listen{
				ArtistName: listen.ArtistName,
				TrackName:  listen.TrackName,
				EndTime:    listen.EndTime,
				MsPlayed:   listen.MsPlayed,
				ID:         id,
			},
		},
	}
	handler.store.AddEvent("user-1", event, version)
}
