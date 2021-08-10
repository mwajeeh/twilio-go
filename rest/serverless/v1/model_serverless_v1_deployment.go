/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ServerlessV1Deployment struct for ServerlessV1Deployment
type ServerlessV1Deployment struct {
	// The SID of the Account that created the Deployment resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Build for the deployment
	BuildSid *string `json:"build_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Deployment resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Deployment resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Environment for the Deployment
	EnvironmentSid *string `json:"environment_sid,omitempty"`
	// The SID of the Service that the Deployment resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Deployment resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Deployment resource
	Url *string `json:"url,omitempty"`
}