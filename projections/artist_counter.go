package projections

import (
	"plotsky.dev/bragi/events"
)

type ArtistCounter struct {
	Artists map[string]int
}

func BuildArtistCounter(all_events []events.Event) ArtistCounter {
	counter := ArtistCounter{
		Artists: make(map[string]int),
	}

	for _, event := range all_events {
		switch event.Data.(type) {
		case events.ListenAdded:
			listen_added := event.Data.(events.ListenAdded)
			if count, ok := counter.Artists[listen_added.Listen.ArtistName]; ok {
				counter.Artists[listen_added.Listen.ArtistName] = count + 1
			} else {
				counter.Artists[listen_added.Listen.ArtistName] = 1
			}
		}
	}

	return counter
}
