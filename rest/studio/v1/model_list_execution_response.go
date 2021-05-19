/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListExecutionResponse struct for ListExecutionResponse
type ListExecutionResponse struct {
	Executions []StudioV1FlowExecution `json:"executions,omitempty"`
	Meta       ListFlowResponseMeta    `json:"meta,omitempty"`
}
