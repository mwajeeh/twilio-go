/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFieldValueResponse struct for ListFieldValueResponse
type ListFieldValueResponse struct {
	FieldValues []AutopilotV1FieldValue   `json:"field_values,omitempty"`
	Meta        ListAssistantResponseMeta `json:"meta,omitempty"`
}
