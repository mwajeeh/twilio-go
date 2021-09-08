/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ServerlessV1BuildStatus struct for ServerlessV1BuildStatus
type ServerlessV1BuildStatus struct {
	// The SID of the Account that created the Build resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Build resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Build resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Build
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Build Status resource
	Url *string `json:"url,omitempty"`
}
