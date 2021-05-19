/*
 * Twilio - Ip_messaging
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

// IpMessagingV2ServiceRole struct for IpMessagingV2ServiceRole
type IpMessagingV2ServiceRole struct {
	AccountSid   *string    `json:"account_sid,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Permissions  *[]string  `json:"permissions,omitempty"`
	ServiceSid   *string    `json:"service_sid,omitempty"`
	Sid          *string    `json:"sid,omitempty"`
	Type         *string    `json:"type,omitempty"`
	Url          *string    `json:"url,omitempty"`
}
