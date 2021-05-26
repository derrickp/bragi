package projections

import (
	"plotsky.dev/bragi/events"
	"plotsky.dev/bragi/listens"
)

type ListenTracker struct {
	listens map[string]listens.Listen
}

func BuildListenTracker(allEvents []events.Event) ListenTracker {
	tracker := ListenTracker{
		listens: make(map[string]listens.Listen),
	}

	for _, event := range allEvents {
		switch event.Data.(type) {
		case events.ListenAdded:
			listen_id := event.Data.(events.ListenAdded).Listen.ID
			tracker.listens[listen_id] = event.Data.(events.ListenAdded).Listen
		}
	}

	return tracker
}

func (tracker ListenTracker) GetListen(listen_id string) (listens.Listen, bool) {
	if listen, ok := tracker.listens[listen_id]; ok {
		return listen, true
	}

	return listens.Listen{}, false
}
