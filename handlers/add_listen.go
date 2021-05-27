package handlers

import (
	"github.com/rs/xid"
	"plotsky.dev/bragi/commands"
	"plotsky.dev/bragi/events"
	"plotsky.dev/bragi/listens"
	"plotsky.dev/bragi/projections"
	"plotsky.dev/bragi/stores"
)

func BuildAddListen(store stores.Store) Handler {
	return Handler{
		store,
	}
}

func (handler Handler) AddListen(command commands.Command) {
	add_spotify_listen, is_listen := command.Data.(commands.AddSpotifyListen)

	if !is_listen {
		return
	}

	spotify_listen := add_spotify_listen.HistoricalListen

	stored_events, version := handler.store.GetEvents(add_spotify_listen.UserId)
	tracker := projections.BuildListenTracker(stored_events)
	id := spotify_listen.EndTime + "-" + spotify_listen.ArtistName + "-" + spotify_listen.TrackName
	println(id)
	_, ok := tracker.GetListen(id)

	println(ok)

	if ok {
		return
	}

	listen_added := events.ListenAdded{
		Listen: listens.FromSpotifyHistoricalListen(id, spotify_listen),
	}

	event := events.Event{
		ID:   xid.New().String(),
		Data: listen_added,
	}
	handler.store.AddEvent(add_spotify_listen.UserId, event, version)
}
