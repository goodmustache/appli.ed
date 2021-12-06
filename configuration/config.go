package configuration

import "encoding/json"

type Formatter string

const (
	FlatFormatter Formatter = "flat"
)

type Config struct {
	WatchList []Watch `json:"watch"`
}

type Watch struct {
	Formatter        Formatter       `json:"formatter"`
	Target           string          `json:"target_config"`
	WatchedLocations []string        `json:"watched_locations"`
	FormatterConfig  json.RawMessage `json:"additional_settings"`
}
