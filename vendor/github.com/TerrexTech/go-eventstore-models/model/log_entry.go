package model

// LogEntry describes the "currently-happening" event.
// Use this to show if everything is going as intended, or reasons for why not.
type LogEntry struct {
	// Description of what happened, can also be an error description.
	Description string `json:"description,omitempty"`
	// ErrorCode is just to inform the kind or classification of error.
	ErrorCode int `json:"errorCode,omitempty"`
	// Level is the severity-level, that is, info, warning, or error.
	Level string `json:"level,omitempty"`
	// EventAction is the action being performed by event corresponding to this log.
	EventAction string `json:"eventAction,omitempty"`
	// EventAction is the service-level action being performed by event
	// corresponding to this log.
	ServiceAction string `json:"serviceAction,omitempty"`
	// ServiceName is the service associated with the log.
	ServiceName string `json:"serviceName,omitempty"`
}
