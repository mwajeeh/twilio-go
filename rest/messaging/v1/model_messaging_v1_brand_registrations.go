/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// MessagingV1BrandRegistrations struct for MessagingV1BrandRegistrations
type MessagingV1BrandRegistrations struct {
	// A2P Messaging Profile Bundle BundleSid
	A2pProfileBundleSid *string `json:"a2p_profile_bundle_sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid
	CustomerProfileBundleSid *string `json:"customer_profile_bundle_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A reason why brand registration has failed
	FailureReason *string `json:"failure_reason,omitempty"`
	// A2P BrandRegistration Sid
	Sid *string `json:"sid,omitempty"`
	// Brand Registration status
	Status *string `json:"status,omitempty"`
	// Campaign Registry (TCR) Brand ID
	TcrId *string `json:"tcr_id,omitempty"`
	// The absolute URL of the Brand Registration
	Url *string `json:"url,omitempty"`
}
