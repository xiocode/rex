package rex

type EventMetadata struct {
	Service     string `json:"service,omitempty"`
	Environment string `json:"env,omitempty"`
	Cluster     string `json:"cluster,omitempty"`
	Host        string `json:"host,omitempty"`
	Release     string `json:"release,omitempty"`
}

type Event struct {
	Ts string `json:"ts,omitempty"`
	EventMetadata
}

type EventBase interface {
	Base() *Event
}
