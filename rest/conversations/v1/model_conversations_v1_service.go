/*
 * Twilio - Conversations
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

// ConversationsV1Service struct for ConversationsV1Service
type ConversationsV1Service struct {
	// The unique ID of the Account responsible for this service.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The human-readable name of this service.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Absolute URLs to access the conversations, users, roles, bindings and configuration of this service.
	Links *map[string]interface{} `json:"links,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// An absolute URL for this service.
	Url *string `json:"url,omitempty"`
}
