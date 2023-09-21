package rocketpkg

type Event struct {
	Topic string
	Tags  []string
}

type EventConf map[EventName]Event

type EventName string
