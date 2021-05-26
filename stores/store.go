package stores

import "plotsky.dev/bragi/events"

type Store interface {
	GetEvents(stream_name string) ([]events.Event, int)
	AddEvent(stream_name string, event events.Event, expected_version int)
}
