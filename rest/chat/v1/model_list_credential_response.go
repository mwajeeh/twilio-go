/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCredentialResponse struct for ListCredentialResponse
type ListCredentialResponse struct {
	Credentials []ChatV1Credential         `json:"credentials,omitempty"`
	Meta        ListCredentialResponseMeta `json:"meta,omitempty"`
}
