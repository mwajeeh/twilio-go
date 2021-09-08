/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1EventType struct for EventsV1EventType
type EventsV1EventType struct {
	// The date this Event Type was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Event Type was updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Event Type description.
	Description *string                 `json:"description,omitempty"`
	Links       *map[string]interface{} `json:"links,omitempty"`
	// The Schema identifier for this Event Type.
	SchemaId *string `json:"schema_id,omitempty"`
	// The Event Type identifier.
	Type *string `json:"type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
