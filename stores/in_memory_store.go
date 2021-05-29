package stores

import (
	"plotsky.dev/bragi/events"
)

type InMemoryStore struct {
	streams map[string][]events.Event
}

func BuildInMemoryStore() InMemoryStore {
	cli := InMemoryStore{
		streams: make(map[string][]events.Event),
	}
	return cli
}

func (store InMemoryStore) AddEvent(stream_name string, event events.Event, expected_version int) {
	stream, ok := store.streams[stream_name]
	if ok {
		store.streams[stream_name] = append(stream, event)
	} else {
		var new_stream []events.Event
		new_stream = append(new_stream, event)
		store.streams[stream_name] = new_stream
	}

}

func (store InMemoryStore) Count(stream_name string) int {
	stream, ok := store.streams[stream_name]
	if ok {
		return len(stream)
	}

	return 0
}

func (store InMemoryStore) GetEvents(stream_name string) ([]events.Event, int) {
	stream, ok := store.streams[stream_name]
	if ok {
		return stream, len(stream)
	}
	var empty []events.Event
	return empty, 0
}
