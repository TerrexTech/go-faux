package model

import (
	"github.com/TerrexTech/uuuid"
)

// Event refers to any interaction with the system, such as inserting or
// retrieving any data.
type Event struct {
	// AggregateID is the ID of aggregate responsible for consuming event.
	AggregateID int8 `cql:"aggregate_id,omitempty" json:"aggregateID,omitempty"`

	// EventAction is the core-action being performed by event.
	// For example, "insert" is EventAction, but "insertUser" is ServiceAction,
	// informing service that a user was inserted.
	EventAction string `cql:"event_action,omitempty" json:"eventAction,omitempty"`

	// ServiceAction is the service-specific Action for the event.
	// For example, "insert" is EventAction, but "insertUser" is ServiceAction,
	// informing service that a user was inserted.
	ServiceAction string `cql:"service_action,omitempty" json:"serviceAction,omitempty"`

	// CorrelationID can be used to "link" events, such as if an event was result of another event.
	// The related events should have cmmon CorrelationIDs, but unique UUIDs.
	// Including CorrelationID will result in inclusion of this ID in any
	// responses generated as per result of event's processing.
	CorrelationID uuuid.UUID `cql:"correlation_id,omitempty" json:"correlationID,omitempty"`

	// Data is the data contained by event.
	Data []byte `cql:"data,omitempty" json:"data,omitempty"`

	// NanoTime is the time in nanoseconds since Unix-epoch to when the event was generated.
	NanoTime int64 `cql:"nano_time,omitempty" json:"nanoTime,omitempty"`

	// UserUUID is the V4-UUID of the user who generated the event.
	UserUUID uuuid.UUID `cql:"user_uuid,omitempty" json:"userUUID,omitempty"`

	// UUID is the V4-UUID unique-indentifier for event.
	UUID uuuid.UUID `cql:"uuid,omitempty" json:"uuid,omitempty"`

	// Version is the version for events as processed for aggregate-projection.
	// This is incremented by the aggregate itself each time it updates its
	// projection.
	Version int64 `cql:"version,omitempty" json:"version,omitempty"`

	// Year bucket is the year in which the event was generated.
	// This is used as the partitioning key.
	YearBucket int16 `cql:"year_bucket,omitempty" json:"yearBucket,omitempty"`
}
