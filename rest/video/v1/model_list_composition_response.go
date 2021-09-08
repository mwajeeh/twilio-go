/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCompositionResponse struct for ListCompositionResponse
type ListCompositionResponse struct {
	Compositions []VideoV1Composition            `json:"compositions,omitempty"`
	Meta         ListCompositionHookResponseMeta `json:"meta,omitempty"`
}
