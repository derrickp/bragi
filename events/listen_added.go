package events

import "plotsky.dev/bragi/listens"

type ListenAdded struct {
	Listen listens.Listen
}
